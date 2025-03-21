import dayjs from "@/lib/dayjs";
import type { OperationState } from "@/lib/grpc-client/buildbarn/buildqueuestate/buildqueuestate";
import themeStyles from "@/theme/theme.module.css";
import { protobufToObjectWithTypeField } from "@/utils/protobufToObject";
import { ExclamationCircleFilled } from "@ant-design/icons";
import { Descriptions, Space, Tag } from "antd";
import Link from "next/link";
import OperationStatusTag from "../OperationStatusTag";
import { operationsStateToActionPageUrl } from "../OperationsGrid/utils";
import PropertyTagList from "../PropertyTagList";

interface Props {
  operation: OperationState;
}

const OperationStateDisplay: React.FC<Props> = ({ operation }) => {
  const invocationMetadata = operation.invocationName?.ids?.map((value) => {
    return JSON.stringify(protobufToObjectWithTypeField(value, false));
  });

  return (
    <Space direction="vertical" size="middle" style={{ display: "flex" }}>
      <Descriptions column={1} size="small" bordered>
        <Descriptions.Item label="Instance name prefix">
          {
            operation.invocationName?.sizeClassQueueName?.platformQueueName
              ?.instanceNamePrefix
          }
        </Descriptions.Item>
        <Descriptions.Item label="Instance name suffix">
          {operation.instanceNameSuffix}
        </Descriptions.Item>
        <Descriptions.Item label="Platform properties">
          <PropertyTagList
            propertyList={
              operation.invocationName?.sizeClassQueueName?.platformQueueName
                ?.platform?.properties
            }
          />
        </Descriptions.Item>
        <Descriptions.Item label="Size class">
          {operation.invocationName?.sizeClassQueueName?.sizeClass}
        </Descriptions.Item>
        <Descriptions.Item label="Invocation IDs">
          <ul>
            {invocationMetadata?.map((value) => (
              <li key={value}>
                <Link
                  href={{
                    pathname: "/operations",
                    query: {
                      filter_invocation_id: value,
                    },
                  }}
                >
                  {value}
                </Link>
              </li>
            ))}
          </ul>
        </Descriptions.Item>
        <Descriptions.Item label="Action digest">
          {operation.actionDigest && (
            <Link
              href={operationsStateToActionPageUrl(operation) || ""}
            >{`${operation.actionDigest.hash}-${operation.actionDigest.sizeBytes}`}</Link>
          )}
        </Descriptions.Item>
        <Descriptions.Item label="Timeout">
          {operation.timeout &&
            `${dayjs(operation.timeout).diff(undefined, "seconds")}s`}
        </Descriptions.Item>
        <Descriptions.Item label="Target ID">
          {operation.targetId}
        </Descriptions.Item>
        <Descriptions.Item label="Priority">
          {operation.priority}
        </Descriptions.Item>
        <Descriptions.Item label="Expected duration">
          {operation.expectedDuration &&
            `${dayjs.duration(Number.parseInt(operation.expectedDuration.seconds), "seconds").asMinutes()}m`}
        </Descriptions.Item>
        <Descriptions.Item label="Age">
          {operation.queuedTimestamp &&
            `${dayjs().diff(operation.queuedTimestamp, "seconds")}s`}
        </Descriptions.Item>
        <Descriptions.Item label="Stage">
          <OperationStatusTag operation={operation} />
          {operation.completed?.status?.message && (
            <Tag
              icon={<ExclamationCircleFilled />}
              color="default"
              className={themeStyles.tag}
            >
              <>Status message: {operation.completed?.status?.message}</>
            </Tag>
          )}
        </Descriptions.Item>
      </Descriptions>
    </Space>
  );
};

export default OperationStateDisplay;
