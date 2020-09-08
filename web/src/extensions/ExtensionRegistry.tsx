import * as H from 'history'
import React, { useEffect, useState, useCallback, useRef } from 'react'
import { PlatformContextProps } from '../../../shared/src/platform/context'
import { Settings, SettingsCascadeProps, SettingsCascadeOrError } from '../../../shared/src/settings/settings'
import { PageTitle } from '../components/PageTitle'
import { eventLogger } from '../tracking/eventLogger'
import { ExtensionsAreaRouteContext } from './ExtensionsArea'
import { ExtensionsList } from './ExtensionsList'
import { ExtensionBanner } from './ExtensionBanner'
import { concat, of, timer } from 'rxjs'
import { debounce, delay, map, switchMap, takeUntil, tap, distinctUntilKeyChanged } from 'rxjs/operators'
import { ConfiguredRegistryExtension, isExtensionEnabled } from '../../../shared/src/extensions/extension'
import { gql } from '../../../shared/src/graphql/graphql'
import { createAggregateError, ErrorLike, isErrorLike } from '../../../shared/src/util/errors'
import { useEventObservable } from '../../../shared/src/util/useObservable'
import {
    RegistryExtensionsResult,
    RegistryExtensionFieldsForList,
    RegistryExtensionsVariables,
} from '../graphql-operations'
import { categorizeExtensionRegistry, CategorizedExtensionRegistry } from './extensions'
import { ExtensionCategory } from '../../../shared/src/schema/extensionSchema'
import { extensionsQuery, isExtensionAdded } from './extension/extension'
import { Link } from 'react-router-dom'
import { Form } from '../components/Form'
import { ExtensionsQueryInputToolbar } from './ExtensionsQueryInputToolbar'

interface Props
    extends Pick<ExtensionsAreaRouteContext, 'authenticatedUser' | 'subject'>,
        PlatformContextProps<'settings' | 'updateSettings' | 'requestGraphQL'>,
        SettingsCascadeProps {
    location: H.Location
    history: H.History
}

const LOADING = 'loading' as const
const URL_QUERY_PARAM = 'query'

export type ExtensionListData = typeof LOADING | (CategorizedExtensionRegistry & { error: string | null }) | ErrorLike

export type ExtensionsEnablement = 'all' | 'enabled' | 'disabled'

const extensionRegistryQuery = gql`
    query RegistryExtensions($query: String, $prioritizeExtensionIDs: [String!]!) {
        extensionRegistry {
            extensions(query: $query, prioritizeExtensionIDs: $prioritizeExtensionIDs) {
                nodes {
                    ...RegistryExtensionFieldsForList
                }
                error
            }
        }
    }
    fragment RegistryExtensionFieldsForList on RegistryExtension {
        id
        publisher {
            __typename
            ... on User {
                id
                username
                displayName
                url
            }
            ... on Org {
                id
                name
                displayName
                url
            }
        }
        extensionID
        extensionIDWithoutRegistry
        name
        manifest {
            raw
            description
        }
        createdAt
        updatedAt
        url
        remoteURL
        registryName
        isLocal
        isWorkInProgress
        viewerCanAdminister
    }
`

/** A page that displays overview information about the available extensions. */
export const ExtensionRegistry: React.FunctionComponent<Props> = props => {
    useEffect(() => eventLogger.logViewEvent('ExtensionsOverview'), [])

    const { history, location, settingsCascade, platformContext, authenticatedUser } = props

    const { current: configuredExtensionCache } = useRef(
        new Map<string, ConfiguredRegistryExtension<RegistryExtensionFieldsForList>>()
    )
    const [query, setQuery] = useState(getQueryFromProps(location))
    const [selectedCategories, setSelectedCategories] = useState<ExtensionCategory[]>([])
    const [enablementFilter, setEnablementFilter] = useState<ExtensionsEnablement>('all')
    const [showMoreExtensions, setShowMoreExtensions] = useState(false)

    /**
     * Note: pass `settingsCascade` instead of making it a dependency to prevent creating
     * new subscriptions when user toggles extensions
     */
    const [nextQueryInput, data] = useEventObservable<
        { query: string; immediate: boolean; settingsCascade: SettingsCascadeOrError<Settings> },
        ExtensionListData
    >(
        useCallback(
            newQueries =>
                newQueries.pipe(
                    distinctUntilKeyChanged('query'),
                    tap(({ query }) => {
                        setQuery(query)

                        history.replace({
                            search: query ? new URLSearchParams({ [URL_QUERY_PARAM]: query }).toString() : '',
                            hash: location.hash,
                        })
                    }),
                    debounce(({ immediate }) => timer(immediate ? 0 : 50)),
                    distinctUntilKeyChanged('query'),
                    switchMap(({ query, immediate, settingsCascade }) => {
                        let viewerConfiguredExtensions: string[] = []
                        if (!isErrorLike(settingsCascade.final)) {
                            if (settingsCascade.final?.extensions) {
                                viewerConfiguredExtensions = Object.keys(settingsCascade.final.extensions)
                            }
                        }
                        const resultOrError = platformContext.requestGraphQL<
                            RegistryExtensionsResult,
                            RegistryExtensionsVariables
                        >({
                            request: extensionRegistryQuery,
                            variables: { query, prioritizeExtensionIDs: viewerConfiguredExtensions },
                            mightContainPrivateInfo: true,
                        })

                        return concat(
                            of(LOADING).pipe(delay(immediate ? 0 : 250), takeUntil(resultOrError)),
                            resultOrError
                        )
                    }),
                    map(resultOrErrorOrLoading => {
                        if (resultOrErrorOrLoading === LOADING) {
                            return resultOrErrorOrLoading
                        }

                        const { data, errors } = resultOrErrorOrLoading

                        if (!data?.extensionRegistry?.extensions) {
                            return createAggregateError(errors)
                        }

                        const { error, nodes } = data.extensionRegistry.extensions

                        return {
                            error,
                            ...categorizeExtensionRegistry(nodes, configuredExtensionCache),
                        }
                    })
                ),
            [platformContext, history, location.hash, configuredExtensionCache]
        )
    )

    const onQueryChangeEvent = useCallback(
        (event: React.FormEvent<HTMLInputElement>) =>
            nextQueryInput({ query: event.currentTarget.value, immediate: false, settingsCascade }),
        [nextQueryInput, settingsCascade]
    )

    const onQueryChangeImmediate = useCallback(
        (query: string) => nextQueryInput({ query, immediate: true, settingsCascade }),
        [nextQueryInput, settingsCascade]
    )

    useEffect(() => {
        // kicks off initial request
        onQueryChangeImmediate(getQueryFromProps(location))
    }, [location, onQueryChangeImmediate])

    const isLoading = !data || data === LOADING

    return (
        <>
            <div className="container">
                <PageTitle title="Extensions" />

                <div className="pt-3">
                    <div className="mb-5">
                        <p>
                            Improve your workflow with code intelligence, test coverage, and other useful information.
                        </p>
                        <Form onSubmit={preventDefault} className="form-inline">
                            <input
                                className="form-control flex-grow-1 mb-2 test-extension-registry-input shadow"
                                type="search"
                                placeholder="Search extensions..."
                                name="query"
                                value={query}
                                onChange={onQueryChangeEvent}
                                autoFocus={true}
                                autoComplete="off"
                                autoCorrect="off"
                                autoCapitalize="off"
                                spellCheck={false}
                            />
                        </Form>
                        <ExtensionsQueryInputToolbar
                            selectedCategories={selectedCategories}
                            onSelectCategories={setSelectedCategories}
                            enablementFilter={enablementFilter}
                            setEnablementFilter={setEnablementFilter}
                        />
                        {!authenticatedUser && (
                            <div className="alert alert-info my-4">
                                <span>An account is required to create and configure extensions. </span>
                                <Link to="/sign-up?returnTo=/extensions">
                                    <span className="alert-link">Register Now!</span>
                                </Link>
                            </div>
                        )}
                        <ExtensionsList
                            {...props}
                            data={data}
                            query={query}
                            enablementFilter={enablementFilter}
                            selectedCategories={selectedCategories}
                            showMoreExtensions={showMoreExtensions}
                        />
                    </div>
                    {!isLoading && !showMoreExtensions && selectedCategories.length === 0 && (
                        <div className="d-flex justify-content-center">
                            <button
                                type="button"
                                className="btn btn-outline-secondary"
                                onClick={() => setShowMoreExtensions(true)}
                            >
                                Show more extensions
                            </button>
                        </div>
                    )}
                </div>
            </div>
            {/* Only show the banner when there are no selected categories and it is not loading */}
            {selectedCategories.length === 0 && !isLoading && (
                <>
                    <hr className="mt-5" />
                    <div className="my-5 row justify-content-center">
                        <div className="mx-auto col-sm-12 col-md-8 col-lg-8 col-xl-6">
                            <ExtensionBanner />
                        </div>
                    </div>
                </>
            )}
        </>
    )
}

function getQueryFromProps(location: H.Location): string {
    const parameters = new URLSearchParams(location.search)
    return parameters.get(URL_QUERY_PARAM) || ''
}

function preventDefault(event: React.FormEvent): void {
    event.preventDefault()
}

/**
 * Applies the query's client-side extensions search keywords #installed, #enabled, and #disabled by filtering
 * {@link registryExtensions}.
 *
 * @internal Exported for testing only.
 */
export function applyExtensionsQuery<X extends { extensionID: string }>(
    query: string,
    settings: Pick<Settings, 'extensions'>,
    registryExtensions: X[]
): X[] {
    const installed = query.includes(extensionsQuery({ installed: true }))
    const enabled = query.includes(extensionsQuery({ enabled: true }))
    const disabled = query.includes(extensionsQuery({ disabled: true }))
    return registryExtensions.filter(
        extension =>
            (!installed || isExtensionAdded(settings, extension.extensionID)) &&
            (!enabled || isExtensionEnabled(settings, extension.extensionID)) &&
            (!disabled ||
                (isExtensionAdded(settings, extension.extensionID) &&
                    !isExtensionEnabled(settings, extension.extensionID)))
    )
}