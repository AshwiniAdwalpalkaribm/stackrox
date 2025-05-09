/* eslint-disable react/no-array-index-key */
import React, { ReactElement } from 'react';
import {
    ActionsColumn,
    ExpandableRowContent,
    IAction,
    Table,
    Tbody,
    Td,
    Th,
    Thead,
    Tr,
} from '@patternfly/react-table';
import {
    Button,
    Flex,
    FlexItem,
    Icon,
    Text,
    TextContent,
    TextVariants,
    Tooltip,
} from '@patternfly/react-core';
import {
    ExclamationCircleIcon,
    ExclamationTriangleIcon,
    MinusIcon,
    PlusIcon,
} from '@patternfly/react-icons';

import { ensureExhaustive } from 'utils/type.utils';
import { BaselineSimulationDiffState, Flow, FlowEntityType } from '../types/flow.type';
import { protocolLabel } from '../utils/flowUtils';

type FlowsTableProps = {
    label: string;
    flows: Flow[];
    numFlows: number;
    expandedRows: string[];
    setExpandedRows: React.Dispatch<React.SetStateAction<string[]>>;
    selectedRows: string[];
    setSelectedRows: React.Dispatch<React.SetStateAction<string[]>>;
    isEditable?: boolean;
    addToBaseline?: (flow: Flow) => void;
    markAsAnomalous?: (flow: Flow) => void;
    isBaselineSimulation?: boolean;
    numExtraneousEgressFlows?: number;
    numExtraneousIngressFlows?: number;
    onSelectFlow: (entityId: string) => void;
};

const columnNames = {
    entity: 'Entity',
    direction: 'Direction',
    // @TODO: This would be a good point to update with i18n translation ability
    portAndProtocol: 'Port / protocol',
};

function getFlowSubtext(flow: Flow): string {
    switch (flow.type) {
        case 'DEPLOYMENT':
            return `in "${flow.namespace}"`;
        case 'CIDR_BLOCK':
        case 'EXTERNAL_ENTITIES':
            return 'External to cluster';
        case 'INTERNAL_ENTITIES':
            return 'Internal to cluster';
        default:
            return ensureExhaustive(flow.type);
    }
}

function getBaselineSimulatedRowStyle(
    baselineSimulationDiffState: BaselineSimulationDiffState | undefined
): React.CSSProperties {
    let customStyle: React.CSSProperties;
    if (baselineSimulationDiffState === 'ADDED') {
        customStyle = { backgroundColor: 'var(--pf-v5-global--palette--green-50)' };
    } else if (baselineSimulationDiffState === 'REMOVED') {
        customStyle = { backgroundColor: 'var(--pf-v5-global--palette--red-50)' };
    } else {
        customStyle = {};
    }
    return customStyle;
}

function ExtraneousFlowsRow({
    isEditable,
    numExtraneousEgressFlows,
    direction,
}: {
    isEditable: boolean;
    numExtraneousEgressFlows: number;
    direction: 'Ingress' | 'Egress';
}) {
    return (
        <Tbody>
            <Tr>
                <Td />
                {isEditable && <Td />}
                <Td dataLabel={columnNames.entity}>
                    <Flex direction={{ default: 'row' }}>
                        <FlexItem>
                            <div>+ {numExtraneousEgressFlows} allowed flows</div>
                            <div>
                                <TextContent>
                                    <Text component={TextVariants.small}>Across this cluster</Text>
                                </TextContent>
                            </div>
                        </FlexItem>
                    </Flex>
                </Td>
                <Td dataLabel={columnNames.direction}>{direction}</Td>
                <Td dataLabel={columnNames.portAndProtocol}>Any / Any</Td>
            </Tr>
        </Tbody>
    );
}

function AnomalousIcon({ type }: { type: FlowEntityType }) {
    if (type === 'CIDR_BLOCK' || type === 'EXTERNAL_ENTITIES') {
        return (
            <Tooltip content={<div>Anomalous external flow</div>}>
                <ExclamationCircleIcon className="pf-v5-u-danger-color-100" />
            </Tooltip>
        );
    }
    return (
        <Tooltip content={<div>Anomalous internal flow</div>}>
            <ExclamationTriangleIcon className="pf-v5-u-warning-color-100" />
        </Tooltip>
    );
}

function FlowsTable({
    label,
    flows,
    numFlows,
    expandedRows,
    setExpandedRows,
    selectedRows,
    setSelectedRows,
    isEditable = false,
    addToBaseline,
    markAsAnomalous,
    isBaselineSimulation = false,
    numExtraneousEgressFlows = 0,
    numExtraneousIngressFlows = 0,
    onSelectFlow,
}: FlowsTableProps): ReactElement {
    // getter functions
    const isRowExpanded = (row: Flow) => expandedRows?.includes(row.id);
    const areAllRowsSelected = selectedRows?.length !== 0 && selectedRows?.length === numFlows;
    const isRowSelected = (row: Flow) => selectedRows?.includes(row.id);

    // setter functions
    const setRowExpanded = (row: Flow, isExpanding = true) =>
        setExpandedRows?.((prevExpanded) => {
            const otherExpandedRows = prevExpanded.filter((r) => r !== row.id);
            return isExpanding ? [...otherExpandedRows, row.id] : otherExpandedRows;
        });
    const setRowSelected = (row: Flow, isSelecting = true) =>
        setSelectedRows?.((prevSelected) => {
            const otherSelectedRows = prevSelected.filter((r) => r !== row.id);
            return isSelecting ? [...otherSelectedRows, row.id] : otherSelectedRows;
        });
    const selectAllRows = (isSelecting = true) => {
        if (isSelecting) {
            const newSelectedRows = flows.reduce((acc, curr) => {
                if (curr.children && curr.children.length !== 0) {
                    return [...acc, ...curr.children.map((child) => child.id)];
                }
                return [...acc, curr.id];
            }, [] as string[]);
            return setSelectedRows?.(newSelectedRows);
        }
        return setSelectedRows?.([]);
    };

    const onSelectFlowHandler = (flow: Flow) => () => {
        onSelectFlow(flow.entityId);
    };

    return (
        <Table aria-label={label} variant="compact">
            <Thead>
                <Tr>
                    <Th>
                        <span className="pf-v5-screen-reader">Row expansion</span>
                    </Th>
                    {isEditable && (
                        <Th
                            select={{
                                onSelect: (_event, isSelecting) => selectAllRows(isSelecting),
                                isSelected: areAllRowsSelected,
                            }}
                        />
                    )}
                    {isBaselineSimulation && <Td />}
                    <Th>{columnNames.entity}</Th>
                    <Th modifier="nowrap">{columnNames.direction}</Th>
                    <Th modifier="nowrap">{columnNames.portAndProtocol}</Th>
                    {isEditable && (
                        <Th>
                            <span className="pf-v5-screen-reader">Row actions</span>
                        </Th>
                    )}
                </Tr>
            </Thead>
            {flows.map((row, rowIndex) => {
                const isExpanded = isRowExpanded(row);
                const rowActions: IAction[] =
                    row.children && !row.children.length
                        ? [
                              row.isAnomalous
                                  ? {
                                        itemKey: 'add-flow-to-baseline',
                                        title: 'Add to baseline',
                                        onClick: () => {
                                            addToBaseline?.(row);
                                        },
                                    }
                                  : {
                                        itemKey: 'mark-flow-as-anomalous',
                                        title: 'Mark as anomalous',
                                        onClick: () => {
                                            markAsAnomalous?.(row);
                                        },
                                    },
                          ]
                        : [];
                const baselineSimulatedRowStyle = getBaselineSimulatedRowStyle(
                    row.baselineSimulationDiffState
                );

                return (
                    <Tbody
                        key={`${row.id}-${rowIndex}`}
                        isExpanded={isExpanded}
                        style={baselineSimulatedRowStyle}
                    >
                        <Tr>
                            <Td
                                expand={
                                    row.children && row.children.length
                                        ? {
                                              rowIndex,
                                              isExpanded,
                                              onToggle: () => setRowExpanded(row, !isExpanded),
                                              expandId: 'flow-expandable',
                                          }
                                        : undefined
                                }
                            />
                            {isEditable && (
                                <Td
                                    select={
                                        row.children && row.children.length === 0
                                            ? {
                                                  rowIndex,
                                                  onSelect: (_event, isSelecting) =>
                                                      setRowSelected(row, isSelecting),
                                                  isSelected: isRowSelected(row),
                                                  isDisabled: !!row.children.length,
                                              }
                                            : undefined
                                    }
                                />
                            )}
                            {isBaselineSimulation && (
                                <Td dataLabel={columnNames.direction}>
                                    {row.baselineSimulationDiffState === 'ADDED' && (
                                        <Tooltip content={<div>Baseline added</div>}>
                                            <Icon size="sm">
                                                <PlusIcon className="pf-v5-u-success-color-200" />
                                            </Icon>
                                        </Tooltip>
                                    )}
                                    {row.baselineSimulationDiffState === 'REMOVED' && (
                                        <Tooltip content={<div>Baseline removed</div>}>
                                            <Icon size="sm">
                                                <MinusIcon className="pf-v5-u-danger-color-200" />
                                            </Icon>
                                        </Tooltip>
                                    )}
                                </Td>
                            )}
                            <Td dataLabel={columnNames.entity}>
                                <Flex direction={{ default: 'row' }}>
                                    <FlexItem>
                                        <div>
                                            <Button
                                                variant="link"
                                                isInline
                                                onClick={onSelectFlowHandler(row)}
                                            >
                                                {row.entity}
                                            </Button>
                                        </div>
                                        <div>
                                            <TextContent>
                                                <Text component={TextVariants.small}>
                                                    {getFlowSubtext(row)}
                                                </Text>
                                            </TextContent>
                                        </div>
                                    </FlexItem>
                                    {row.isAnomalous && (
                                        <FlexItem>
                                            <AnomalousIcon type={row.type} />
                                        </FlexItem>
                                    )}
                                </Flex>
                            </Td>
                            <Td dataLabel={columnNames.direction}>{row.direction}</Td>
                            <Td dataLabel={columnNames.portAndProtocol}>
                                {row.port} / {protocolLabel[row.protocol]}
                            </Td>
                            {isEditable && (
                                <Td isActionCell>
                                    {row.children && !row.children.length && (
                                        <ActionsColumn items={rowActions} />
                                    )}
                                </Td>
                            )}
                        </Tr>
                        {isExpanded &&
                            row.children &&
                            row.children.map((child, index) => {
                                const childActions: IAction[] = [
                                    child.isAnomalous
                                        ? {
                                              itemKey: 'add-flow-to-baseline',
                                              title: 'Add to baseline',
                                              onClick: () => {
                                                  addToBaseline?.(child);
                                              },
                                          }
                                        : {
                                              itemKey: 'mark-flow-as-anomalous',
                                              title: 'Mark as anomalous',
                                              onClick: () => {
                                                  markAsAnomalous?.(child);
                                              },
                                          },
                                ];

                                return (
                                    <Tr key={`${child.id}-${index}`} isExpanded={isExpanded}>
                                        <Td />
                                        {isEditable && (
                                            <Td
                                                select={{
                                                    rowIndex,
                                                    onSelect: (_event, isSelecting) =>
                                                        setRowSelected(child, isSelecting),
                                                    isSelected: isRowSelected(child),
                                                }}
                                            />
                                        )}
                                        <Td>
                                            <ExpandableRowContent>
                                                <Flex direction={{ default: 'row' }}>
                                                    <FlexItem>
                                                        <Button
                                                            variant="link"
                                                            isInline
                                                            onClick={onSelectFlowHandler(child)}
                                                        >
                                                            {child.entity}
                                                        </Button>
                                                    </FlexItem>
                                                    {child.isAnomalous && (
                                                        <FlexItem>
                                                            <AnomalousIcon type={child.type} />
                                                        </FlexItem>
                                                    )}
                                                </Flex>
                                            </ExpandableRowContent>
                                        </Td>
                                        <Td>
                                            <ExpandableRowContent>
                                                {child.direction}
                                            </ExpandableRowContent>
                                        </Td>
                                        <Td>
                                            <ExpandableRowContent>
                                                {child.port} / {protocolLabel[child.protocol]}
                                            </ExpandableRowContent>
                                        </Td>
                                        {isEditable && (
                                            <Td isActionCell>
                                                <ActionsColumn items={childActions} />
                                            </Td>
                                        )}
                                    </Tr>
                                );
                            })}
                    </Tbody>
                );
            })}
            {numExtraneousEgressFlows > 0 && (
                <ExtraneousFlowsRow
                    isEditable
                    numExtraneousEgressFlows={numExtraneousEgressFlows}
                    direction="Egress"
                />
            )}
            {numExtraneousIngressFlows > 0 && (
                <ExtraneousFlowsRow
                    isEditable
                    numExtraneousEgressFlows={numExtraneousIngressFlows}
                    direction="Ingress"
                />
            )}
        </Table>
    );
}

export default FlowsTable;
