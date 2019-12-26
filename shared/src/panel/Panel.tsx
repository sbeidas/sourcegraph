import * as H from 'history'
import CloseIcon from 'mdi-react/CloseIcon'
import * as React from 'react'
import { Observable, Subscription } from 'rxjs'
import { map } from 'rxjs/operators'
import { ActionsNavItems } from '../actions/ActionsNavItems'
import { PanelViewWithComponent, ViewProviderRegistrationOptions } from '../api/client/services/view'
import { ContributableMenu, ContributableViewContainer } from '../api/protocol/contribution'
import { ActivationProps } from '../components/activation/Activation'
import { FetchFileCtx } from '../components/CodeExcerpt'
import { Resizable } from '../components/Resizable'
import { Section } from '../components/sections/Sections'
import {
    TabbedSectionsWithURLViewStatePersistence,
    Spacer,
    currentSectionForTabbedSectionsWithURLViewStatePersistence,
    urlForSectionForTabbedSectionsWithURLViewStatePersistence,
} from '../components/sections/tabbed/TabbedSections'
import { ExtensionsControllerProps } from '../extensions/controller'
import { PlatformContextProps } from '../platform/context'
import { SettingsCascadeProps } from '../settings/settings'
import { TelemetryProps } from '../telemetry/telemetryService'
import { ThemeProps } from '../theme'
import { EmptyPanelView } from './views/EmptyPanelView'
import { PanelView } from './views/PanelView'

interface Props
    extends ExtensionsControllerProps,
        PlatformContextProps,
        SettingsCascadeProps,
        ActivationProps,
        TelemetryProps,
        ThemeProps {
    location: H.Location
    history: H.History
    repoName?: string
    fetchHighlightedFileLines: (ctx: FetchFileCtx, force?: boolean) => Observable<string[]>
}

interface State {
    /** Panel views contributed by extensions. */
    panelViews?: (PanelViewWithComponent & Pick<ViewProviderRegistrationOptions, 'id'>)[] | null
}

/**
 * A tab and corresponding content to display in the panel.
 */
interface PanelItem extends Section<string> {
    /**
     * Controls the relative order of panel items. The items are laid out from highest priority (at the beginning)
     * to lowest priority (at the end). The default is 0.
     */
    priority: number

    /** The content element to display when the tab is active. */
    element: JSX.Element

    /**
     * Whether this panel contains a list of locations (from a location provider). This value is
     * exposed to contributions as `panel.activeView.hasLocations`. It is true if there is a
     * location provider (even if the result set is empty).
     */
    hasLocations?: boolean
}

/**
 * The panel, which is a tabbed component with contextual information. Components rendering the panel should
 * generally use ResizablePanel, not Panel.
 *
 * Other components can contribute panel items to the panel.
 */
export class Panel extends React.PureComponent<Props, State> {
    public state: State = {}

    private subscriptions = new Subscription()

    public componentDidMount(): void {
        this.subscriptions.add(
            this.props.extensionsController.services.views
                .getViews(ContributableViewContainer.Panel)
                .pipe(map(panelViews => ({ panelViews })))
                .subscribe(stateUpdate => this.setState(stateUpdate))
        )
    }

    public componentWillUnmount(): void {
        this.subscriptions.unsubscribe()
    }

    public render(): JSX.Element | null {
        const items = this.state.panelViews
            ? this.state.panelViews
                  .map(
                      (panelView): PanelItem => ({
                          label: panelView.title,
                          id: panelView.id,
                          priority: panelView.priority,
                          element: <PanelView {...this.props} panelView={panelView} />,
                          hasLocations: !!panelView.locationProvider,
                      })
                  )
                  .sort(byPriority)
            : []

        const hasTabs = items.length > 0
        const activePanelViewID = currentSectionForTabbedSectionsWithURLViewStatePersistence(items, this.props.location)
        const activePanelView = items.find(item => item.id === activePanelViewID)

        return (
            <div className="panel">
                {hasTabs ? (
                    <TabbedSectionsWithURLViewStatePersistence
                        sections={items}
                        navbarEndFragment={
                            <>
                                <Spacer />
                                <button
                                    type="button"
                                    onClick={this.onDismiss}
                                    className="btn btn-icon tabbed-sections__navbar-close-button tabbed-sections__navbar-end-fragment-other-element"
                                    data-tooltip="Close"
                                >
                                    <CloseIcon className="icon-inline" />
                                </button>
                            </>
                        }
                        toolbarFragment={
                            <ActionsNavItems
                                {...this.props}
                                // TODO remove references to Bootstrap from shared, get class name from prop
                                // This is okay for now because the Panel is currently only used in the webapp
                                listClass="nav w-100 justify-content-end"
                                actionItemClass="nav-link"
                                actionItemIconClass="icon-inline"
                                menu={ContributableMenu.PanelToolbar}
                                scope={
                                    activePanelViewID !== undefined
                                        ? {
                                              type: 'panelView',
                                              id: activePanelViewID,
                                              hasLocations: Boolean(activePanelView?.hasLocations),
                                          }
                                        : undefined
                                }
                                wrapInList={true}
                            />
                        }
                        className="panel__tabs"
                        navbarItemClassName="tabbed-sections__navbar-tab--h5like"
                        location={this.props.location}
                    >
                        {items?.map(({ id, element }) => React.cloneElement(element, { key: id }))}
                    </TabbedSectionsWithURLViewStatePersistence>
                ) : (
                    <EmptyPanelView />
                )}
            </div>
        )
    }

    private onDismiss = (): void =>
        this.props.history.push(urlForSectionForTabbedSectionsWithURLViewStatePersistence(null, this.props.location))
}

function byPriority(a: { priority: number }, b: { priority: number }): number {
    return b.priority - a.priority
}

/** A wrapper around Panel that makes it resizable. */
export const ResizablePanel: React.FunctionComponent<Props> = props => (
    <Resizable
        className="panel--resizable"
        handlePosition="top"
        defaultSize={350}
        storageKey="panel-size"
        element={<Panel {...props} />}
    />
)
