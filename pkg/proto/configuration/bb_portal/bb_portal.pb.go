// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: pkg/proto/configuration/bb_portal/bb_portal.proto

package bb_portal

import (
	auth "github.com/buildbarn/bb-storage/pkg/proto/configuration/auth"
	blobstore "github.com/buildbarn/bb-storage/pkg/proto/configuration/blobstore"
	global "github.com/buildbarn/bb-storage/pkg/proto/configuration/global"
	grpc "github.com/buildbarn/bb-storage/pkg/proto/configuration/grpc"
	http "github.com/buildbarn/bb-storage/pkg/proto/configuration/http"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostgresSource struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ConnectionString string                 `protobuf:"bytes,1,opt,name=connection_string,json=connectionString,proto3" json:"connection_string,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *PostgresSource) Reset() {
	*x = PostgresSource{}
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostgresSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresSource) ProtoMessage() {}

func (x *PostgresSource) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresSource.ProtoReflect.Descriptor instead.
func (*PostgresSource) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescGZIP(), []int{0}
}

func (x *PostgresSource) GetConnectionString() string {
	if x != nil {
		return x.ConnectionString
	}
	return ""
}

type SqliteSource struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ConnectionString string                 `protobuf:"bytes,1,opt,name=connection_string,json=connectionString,proto3" json:"connection_string,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *SqliteSource) Reset() {
	*x = SqliteSource{}
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SqliteSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqliteSource) ProtoMessage() {}

func (x *SqliteSource) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqliteSource.ProtoReflect.Descriptor instead.
func (*SqliteSource) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescGZIP(), []int{1}
}

func (x *SqliteSource) GetConnectionString() string {
	if x != nil {
		return x.ConnectionString
	}
	return ""
}

type Database struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Source:
	//
	//	*Database_Sqlite
	//	*Database_Postgres
	Source        isDatabase_Source `protobuf_oneof:"source"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Database) Reset() {
	*x = Database{}
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescGZIP(), []int{2}
}

func (x *Database) GetSource() isDatabase_Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Database) GetSqlite() *SqliteSource {
	if x != nil {
		if x, ok := x.Source.(*Database_Sqlite); ok {
			return x.Sqlite
		}
	}
	return nil
}

func (x *Database) GetPostgres() *PostgresSource {
	if x != nil {
		if x, ok := x.Source.(*Database_Postgres); ok {
			return x.Postgres
		}
	}
	return nil
}

type isDatabase_Source interface {
	isDatabase_Source()
}

type Database_Sqlite struct {
	Sqlite *SqliteSource `protobuf:"bytes,1,opt,name=sqlite,proto3,oneof"`
}

type Database_Postgres struct {
	Postgres *PostgresSource `protobuf:"bytes,2,opt,name=postgres,proto3,oneof"`
}

func (*Database_Sqlite) isDatabase_Source() {}

func (*Database_Postgres) isDatabase_Source() {}

type BuildEventStreamService struct {
	state                   protoimpl.MessageState      `protogen:"open.v1"`
	GrpcServers             []*grpc.ServerConfiguration `protobuf:"bytes,1,rep,name=grpc_servers,json=grpcServers,proto3" json:"grpc_servers,omitempty"`
	Database                *Database                   `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	BlobArchiveFolder       string                      `protobuf:"bytes,3,opt,name=blob_archive_folder,json=blobArchiveFolder,proto3" json:"blob_archive_folder,omitempty"`
	EnableBepFileUpload     bool                        `protobuf:"varint,4,opt,name=enable_bep_file_upload,json=enableBepFileUpload,proto3" json:"enable_bep_file_upload,omitempty"`
	EnableGraphqlPlayground bool                        `protobuf:"varint,5,opt,name=enable_graphql_playground,json=enableGraphqlPlayground,proto3" json:"enable_graphql_playground,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *BuildEventStreamService) Reset() {
	*x = BuildEventStreamService{}
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuildEventStreamService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildEventStreamService) ProtoMessage() {}

func (x *BuildEventStreamService) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildEventStreamService.ProtoReflect.Descriptor instead.
func (*BuildEventStreamService) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescGZIP(), []int{3}
}

func (x *BuildEventStreamService) GetGrpcServers() []*grpc.ServerConfiguration {
	if x != nil {
		return x.GrpcServers
	}
	return nil
}

func (x *BuildEventStreamService) GetDatabase() *Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *BuildEventStreamService) GetBlobArchiveFolder() string {
	if x != nil {
		return x.BlobArchiveFolder
	}
	return ""
}

func (x *BuildEventStreamService) GetEnableBepFileUpload() bool {
	if x != nil {
		return x.EnableBepFileUpload
	}
	return false
}

func (x *BuildEventStreamService) GetEnableGraphqlPlayground() bool {
	if x != nil {
		return x.EnableGraphqlPlayground
	}
	return false
}

type BrowserService struct {
	state                     protoimpl.MessageState             `protogen:"open.v1"`
	ContentAddressableStorage *blobstore.BlobAccessConfiguration `protobuf:"bytes,1,opt,name=content_addressable_storage,json=contentAddressableStorage,proto3" json:"content_addressable_storage,omitempty"`
	ActionCache               *blobstore.BlobAccessConfiguration `protobuf:"bytes,2,opt,name=action_cache,json=actionCache,proto3" json:"action_cache,omitempty"`
	InitialSizeClassCache     *blobstore.BlobAccessConfiguration `protobuf:"bytes,3,opt,name=initial_size_class_cache,json=initialSizeClassCache,proto3" json:"initial_size_class_cache,omitempty"`
	FileSystemAccessCache     *blobstore.BlobAccessConfiguration `protobuf:"bytes,4,opt,name=file_system_access_cache,json=fileSystemAccessCache,proto3" json:"file_system_access_cache,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *BrowserService) Reset() {
	*x = BrowserService{}
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrowserService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowserService) ProtoMessage() {}

func (x *BrowserService) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowserService.ProtoReflect.Descriptor instead.
func (*BrowserService) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescGZIP(), []int{4}
}

func (x *BrowserService) GetContentAddressableStorage() *blobstore.BlobAccessConfiguration {
	if x != nil {
		return x.ContentAddressableStorage
	}
	return nil
}

func (x *BrowserService) GetActionCache() *blobstore.BlobAccessConfiguration {
	if x != nil {
		return x.ActionCache
	}
	return nil
}

func (x *BrowserService) GetInitialSizeClassCache() *blobstore.BlobAccessConfiguration {
	if x != nil {
		return x.InitialSizeClassCache
	}
	return nil
}

func (x *BrowserService) GetFileSystemAccessCache() *blobstore.BlobAccessConfiguration {
	if x != nil {
		return x.FileSystemAccessCache
	}
	return nil
}

type SchedulerService struct {
	state                    protoimpl.MessageState        `protogen:"open.v1"`
	BuildQueueStateClient    *grpc.ClientConfiguration     `protobuf:"bytes,1,opt,name=build_queue_state_client,json=buildQueueStateClient,proto3" json:"build_queue_state_client,omitempty"`
	KillOperationsAuthorizer *auth.AuthorizerConfiguration `protobuf:"bytes,2,opt,name=kill_operations_authorizer,json=killOperationsAuthorizer,proto3" json:"kill_operations_authorizer,omitempty"`
	ListOperationsPageSize   uint32                        `protobuf:"varint,3,opt,name=listOperationsPageSize,proto3" json:"listOperationsPageSize,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *SchedulerService) Reset() {
	*x = SchedulerService{}
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchedulerService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerService) ProtoMessage() {}

func (x *SchedulerService) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerService.ProtoReflect.Descriptor instead.
func (*SchedulerService) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescGZIP(), []int{5}
}

func (x *SchedulerService) GetBuildQueueStateClient() *grpc.ClientConfiguration {
	if x != nil {
		return x.BuildQueueStateClient
	}
	return nil
}

func (x *SchedulerService) GetKillOperationsAuthorizer() *auth.AuthorizerConfiguration {
	if x != nil {
		return x.KillOperationsAuthorizer
	}
	return nil
}

func (x *SchedulerService) GetListOperationsPageSize() uint32 {
	if x != nil {
		return x.ListOperationsPageSize
	}
	return 0
}

type ApplicationConfiguration struct {
	state                         protoimpl.MessageState        `protogen:"open.v1"`
	HttpServers                   []*http.ServerConfiguration   `protobuf:"bytes,1,rep,name=http_servers,json=httpServers,proto3" json:"http_servers,omitempty"`
	Global                        *global.Configuration         `protobuf:"bytes,2,opt,name=global,proto3" json:"global,omitempty"`
	BesServiceConfiguration       *BuildEventStreamService      `protobuf:"bytes,3,opt,name=bes_service_configuration,json=besServiceConfiguration,proto3" json:"bes_service_configuration,omitempty"`
	BrowserServiceConfiguration   *BrowserService               `protobuf:"bytes,4,opt,name=browser_service_configuration,json=browserServiceConfiguration,proto3" json:"browser_service_configuration,omitempty"`
	SchedulerServiceConfiguration *SchedulerService             `protobuf:"bytes,5,opt,name=scheduler_service_configuration,json=schedulerServiceConfiguration,proto3" json:"scheduler_service_configuration,omitempty"`
	MaximumMessageSizeBytes       int64                         `protobuf:"varint,6,opt,name=maximum_message_size_bytes,json=maximumMessageSizeBytes,proto3" json:"maximum_message_size_bytes,omitempty"`
	InstanceNameAuthorizer        *auth.AuthorizerConfiguration `protobuf:"bytes,7,opt,name=instance_name_authorizer,json=instanceNameAuthorizer,proto3" json:"instance_name_authorizer,omitempty"`
	FrontendProxyUrl              string                        `protobuf:"bytes,8,opt,name=frontend_proxy_url,json=frontendProxyUrl,proto3" json:"frontend_proxy_url,omitempty"`
	AllowedOrigins                []string                      `protobuf:"bytes,9,rep,name=allowed_origins,json=allowedOrigins,proto3" json:"allowed_origins,omitempty"`
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *ApplicationConfiguration) Reset() {
	*x = ApplicationConfiguration{}
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplicationConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationConfiguration) ProtoMessage() {}

func (x *ApplicationConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationConfiguration.ProtoReflect.Descriptor instead.
func (*ApplicationConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescGZIP(), []int{6}
}

func (x *ApplicationConfiguration) GetHttpServers() []*http.ServerConfiguration {
	if x != nil {
		return x.HttpServers
	}
	return nil
}

func (x *ApplicationConfiguration) GetGlobal() *global.Configuration {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *ApplicationConfiguration) GetBesServiceConfiguration() *BuildEventStreamService {
	if x != nil {
		return x.BesServiceConfiguration
	}
	return nil
}

func (x *ApplicationConfiguration) GetBrowserServiceConfiguration() *BrowserService {
	if x != nil {
		return x.BrowserServiceConfiguration
	}
	return nil
}

func (x *ApplicationConfiguration) GetSchedulerServiceConfiguration() *SchedulerService {
	if x != nil {
		return x.SchedulerServiceConfiguration
	}
	return nil
}

func (x *ApplicationConfiguration) GetMaximumMessageSizeBytes() int64 {
	if x != nil {
		return x.MaximumMessageSizeBytes
	}
	return 0
}

func (x *ApplicationConfiguration) GetInstanceNameAuthorizer() *auth.AuthorizerConfiguration {
	if x != nil {
		return x.InstanceNameAuthorizer
	}
	return nil
}

func (x *ApplicationConfiguration) GetFrontendProxyUrl() string {
	if x != nil {
		return x.FrontendProxyUrl
	}
	return ""
}

func (x *ApplicationConfiguration) GetAllowedOrigins() []string {
	if x != nil {
		return x.AllowedOrigins
	}
	return nil
}

var File_pkg_proto_configuration_bb_portal_bb_portal_proto protoreflect.FileDescriptor

const file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDesc = "" +
	"\n" +
	"1pkg/proto/configuration/bb_portal/bb_portal.proto\x12!buildbarn.configuration.bb_portal\x1a'pkg/proto/configuration/auth/auth.proto\x1a1pkg/proto/configuration/blobstore/blobstore.proto\x1a+pkg/proto/configuration/global/global.proto\x1a'pkg/proto/configuration/grpc/grpc.proto\x1a'pkg/proto/configuration/http/http.proto\"=\n" +
	"\x0ePostgresSource\x12+\n" +
	"\x11connection_string\x18\x01 \x01(\tR\x10connectionString\";\n" +
	"\fSqliteSource\x12+\n" +
	"\x11connection_string\x18\x01 \x01(\tR\x10connectionString\"\xb0\x01\n" +
	"\bDatabase\x12I\n" +
	"\x06sqlite\x18\x01 \x01(\v2/.buildbarn.configuration.bb_portal.SqliteSourceH\x00R\x06sqlite\x12O\n" +
	"\bpostgres\x18\x02 \x01(\v21.buildbarn.configuration.bb_portal.PostgresSourceH\x00R\bpostgresB\b\n" +
	"\x06source\"\xd9\x02\n" +
	"\x17BuildEventStreamService\x12T\n" +
	"\fgrpc_servers\x18\x01 \x03(\v21.buildbarn.configuration.grpc.ServerConfigurationR\vgrpcServers\x12G\n" +
	"\bdatabase\x18\x02 \x01(\v2+.buildbarn.configuration.bb_portal.DatabaseR\bdatabase\x12.\n" +
	"\x13blob_archive_folder\x18\x03 \x01(\tR\x11blobArchiveFolder\x123\n" +
	"\x16enable_bep_file_upload\x18\x04 \x01(\bR\x13enableBepFileUpload\x12:\n" +
	"\x19enable_graphql_playground\x18\x05 \x01(\bR\x17enableGraphqlPlayground\"\xd5\x03\n" +
	"\x0eBrowserService\x12z\n" +
	"\x1bcontent_addressable_storage\x18\x01 \x01(\v2:.buildbarn.configuration.blobstore.BlobAccessConfigurationR\x19contentAddressableStorage\x12]\n" +
	"\faction_cache\x18\x02 \x01(\v2:.buildbarn.configuration.blobstore.BlobAccessConfigurationR\vactionCache\x12s\n" +
	"\x18initial_size_class_cache\x18\x03 \x01(\v2:.buildbarn.configuration.blobstore.BlobAccessConfigurationR\x15initialSizeClassCache\x12s\n" +
	"\x18file_system_access_cache\x18\x04 \x01(\v2:.buildbarn.configuration.blobstore.BlobAccessConfigurationR\x15fileSystemAccessCache\"\xab\x02\n" +
	"\x10SchedulerService\x12j\n" +
	"\x18build_queue_state_client\x18\x01 \x01(\v21.buildbarn.configuration.grpc.ClientConfigurationR\x15buildQueueStateClient\x12s\n" +
	"\x1akill_operations_authorizer\x18\x02 \x01(\v25.buildbarn.configuration.auth.AuthorizerConfigurationR\x18killOperationsAuthorizer\x126\n" +
	"\x16listOperationsPageSize\x18\x03 \x01(\rR\x16listOperationsPageSize\"\xa8\x06\n" +
	"\x18ApplicationConfiguration\x12T\n" +
	"\fhttp_servers\x18\x01 \x03(\v21.buildbarn.configuration.http.ServerConfigurationR\vhttpServers\x12E\n" +
	"\x06global\x18\x02 \x01(\v2-.buildbarn.configuration.global.ConfigurationR\x06global\x12v\n" +
	"\x19bes_service_configuration\x18\x03 \x01(\v2:.buildbarn.configuration.bb_portal.BuildEventStreamServiceR\x17besServiceConfiguration\x12u\n" +
	"\x1dbrowser_service_configuration\x18\x04 \x01(\v21.buildbarn.configuration.bb_portal.BrowserServiceR\x1bbrowserServiceConfiguration\x12{\n" +
	"\x1fscheduler_service_configuration\x18\x05 \x01(\v23.buildbarn.configuration.bb_portal.SchedulerServiceR\x1dschedulerServiceConfiguration\x12;\n" +
	"\x1amaximum_message_size_bytes\x18\x06 \x01(\x03R\x17maximumMessageSizeBytes\x12o\n" +
	"\x18instance_name_authorizer\x18\a \x01(\v25.buildbarn.configuration.auth.AuthorizerConfigurationR\x16instanceNameAuthorizer\x12,\n" +
	"\x12frontend_proxy_url\x18\b \x01(\tR\x10frontendProxyUrl\x12'\n" +
	"\x0fallowed_origins\x18\t \x03(\tR\x0eallowedOriginsBBZ@github.com/buildbarn/bb-portal/pkg/proto/configuration/bb_portalb\x06proto3"

var (
	file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescData []byte
)

func file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDesc), len(file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDesc)))
	})
	return file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDescData
}

var file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_configuration_bb_portal_bb_portal_proto_goTypes = []any{
	(*PostgresSource)(nil),                    // 0: buildbarn.configuration.bb_portal.PostgresSource
	(*SqliteSource)(nil),                      // 1: buildbarn.configuration.bb_portal.SqliteSource
	(*Database)(nil),                          // 2: buildbarn.configuration.bb_portal.Database
	(*BuildEventStreamService)(nil),           // 3: buildbarn.configuration.bb_portal.BuildEventStreamService
	(*BrowserService)(nil),                    // 4: buildbarn.configuration.bb_portal.BrowserService
	(*SchedulerService)(nil),                  // 5: buildbarn.configuration.bb_portal.SchedulerService
	(*ApplicationConfiguration)(nil),          // 6: buildbarn.configuration.bb_portal.ApplicationConfiguration
	(*grpc.ServerConfiguration)(nil),          // 7: buildbarn.configuration.grpc.ServerConfiguration
	(*blobstore.BlobAccessConfiguration)(nil), // 8: buildbarn.configuration.blobstore.BlobAccessConfiguration
	(*grpc.ClientConfiguration)(nil),          // 9: buildbarn.configuration.grpc.ClientConfiguration
	(*auth.AuthorizerConfiguration)(nil),      // 10: buildbarn.configuration.auth.AuthorizerConfiguration
	(*http.ServerConfiguration)(nil),          // 11: buildbarn.configuration.http.ServerConfiguration
	(*global.Configuration)(nil),              // 12: buildbarn.configuration.global.Configuration
}
var file_pkg_proto_configuration_bb_portal_bb_portal_proto_depIdxs = []int32{
	1,  // 0: buildbarn.configuration.bb_portal.Database.sqlite:type_name -> buildbarn.configuration.bb_portal.SqliteSource
	0,  // 1: buildbarn.configuration.bb_portal.Database.postgres:type_name -> buildbarn.configuration.bb_portal.PostgresSource
	7,  // 2: buildbarn.configuration.bb_portal.BuildEventStreamService.grpc_servers:type_name -> buildbarn.configuration.grpc.ServerConfiguration
	2,  // 3: buildbarn.configuration.bb_portal.BuildEventStreamService.database:type_name -> buildbarn.configuration.bb_portal.Database
	8,  // 4: buildbarn.configuration.bb_portal.BrowserService.content_addressable_storage:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	8,  // 5: buildbarn.configuration.bb_portal.BrowserService.action_cache:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	8,  // 6: buildbarn.configuration.bb_portal.BrowserService.initial_size_class_cache:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	8,  // 7: buildbarn.configuration.bb_portal.BrowserService.file_system_access_cache:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	9,  // 8: buildbarn.configuration.bb_portal.SchedulerService.build_queue_state_client:type_name -> buildbarn.configuration.grpc.ClientConfiguration
	10, // 9: buildbarn.configuration.bb_portal.SchedulerService.kill_operations_authorizer:type_name -> buildbarn.configuration.auth.AuthorizerConfiguration
	11, // 10: buildbarn.configuration.bb_portal.ApplicationConfiguration.http_servers:type_name -> buildbarn.configuration.http.ServerConfiguration
	12, // 11: buildbarn.configuration.bb_portal.ApplicationConfiguration.global:type_name -> buildbarn.configuration.global.Configuration
	3,  // 12: buildbarn.configuration.bb_portal.ApplicationConfiguration.bes_service_configuration:type_name -> buildbarn.configuration.bb_portal.BuildEventStreamService
	4,  // 13: buildbarn.configuration.bb_portal.ApplicationConfiguration.browser_service_configuration:type_name -> buildbarn.configuration.bb_portal.BrowserService
	5,  // 14: buildbarn.configuration.bb_portal.ApplicationConfiguration.scheduler_service_configuration:type_name -> buildbarn.configuration.bb_portal.SchedulerService
	10, // 15: buildbarn.configuration.bb_portal.ApplicationConfiguration.instance_name_authorizer:type_name -> buildbarn.configuration.auth.AuthorizerConfiguration
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_bb_portal_bb_portal_proto_init() }
func file_pkg_proto_configuration_bb_portal_bb_portal_proto_init() {
	if File_pkg_proto_configuration_bb_portal_bb_portal_proto != nil {
		return
	}
	file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes[2].OneofWrappers = []any{
		(*Database_Sqlite)(nil),
		(*Database_Postgres)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDesc), len(file_pkg_proto_configuration_bb_portal_bb_portal_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_bb_portal_bb_portal_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_bb_portal_bb_portal_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_bb_portal_bb_portal_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_bb_portal_bb_portal_proto = out.File
	file_pkg_proto_configuration_bb_portal_bb_portal_proto_goTypes = nil
	file_pkg_proto_configuration_bb_portal_bb_portal_proto_depIdxs = nil
}
