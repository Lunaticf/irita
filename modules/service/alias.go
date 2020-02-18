// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/bianjieai/irita/modules/service/internal/keeper
// ALIASGEN: github.com/bianjieai/irita/modules/service/internal/types
package service

import (
	"github.com/bianjieai/irita/modules/service/internal/keeper"
	"github.com/bianjieai/irita/modules/service/internal/types"
)

const (
	Global                  = types.Global
	Local                   = types.Local
	NoPrivacy               = types.NoPrivacy
	PubKeyEncryption        = types.PubKeyEncryption
	OffChainCached          = types.OffChainCached
	NoCached                = types.NoCached
	Unicast                 = types.Unicast
	Multicast               = types.Multicast
	EventTypeRequestSvc     = types.EventTypeRequestSvc
	EventTypeRespondSvc     = types.EventTypeRespondSvc
	EventTypeSvcCallTimeout = types.EventTypeSvcCallTimeout
	AttributeKeyProvider    = types.AttributeKeyProvider
	AttributeKeyConsumer    = types.AttributeKeyConsumer
	AttributeKeyRequestID   = types.AttributeKeyRequestID
	AttributeKeyServiceFee  = types.AttributeKeyServiceFee
	AttributeKeySlashCoins  = types.AttributeKeySlashCoins
	AttributeValueCategory  = types.AttributeValueCategory
	ModuleName              = types.ModuleName
	StoreKey                = types.StoreKey
	QuerierRoute            = types.QuerierRoute
	RouterKey               = types.RouterKey
	DefaultParamspace       = types.DefaultParamspace
	DepositAccName          = types.DepositAccName
	RequestAccName          = types.RequestAccName
	TaxAccName              = types.TaxAccName
	MetricsSubsystem        = types.MetricsSubsystem
	TypeMsgSvcDef           = types.TypeMsgSvcDef
	TypeMsgSvcBind          = types.TypeMsgSvcBind
	TypeMsgSvcBindingUpdate = types.TypeMsgSvcBindingUpdate
	TypeMsgSvcDisable       = types.TypeMsgSvcDisable
	TypeMsgSvcEnable        = types.TypeMsgSvcEnable
	TypeMsgSvcRefundDeposit = types.TypeMsgSvcRefundDeposit
	TypeMsgSvcRequest       = types.TypeMsgSvcRequest
	TypeMsgSvcResponse      = types.TypeMsgSvcResponse
	TypeMsgSvcRefundFees    = types.TypeMsgSvcRefundFees
	TypeMsgSvcWithdrawFees  = types.TypeMsgSvcWithdrawFees
	TypeMsgSvcWithdrawTax   = types.TypeMsgSvcWithdrawTax
	MaxNameLength           = types.MaxNameLength
	MaxChainIDLength        = types.MaxChainIDLength
	MaxDescriptionLength    = types.MaxDescriptionLength
	MaxTagCount             = types.MaxTagCount
	MaxTagLength            = types.MaxTagLength
	QueryDefinition         = types.QueryDefinition
	QueryBinding            = types.QueryBinding
	QueryBindings           = types.QueryBindings
	QueryRequests           = types.QueryRequests
	QueryResponse           = types.QueryResponse
	QueryFees               = types.QueryFees
)

var (
	// functions aliases
	NewKeeper                            = keeper.NewKeeper
	ParamKeyTable                        = keeper.ParamKeyTable
	NewQuerier                           = keeper.NewQuerier
	NewSvcBinding                        = types.NewSvcBinding
	SvcBindingEqual                      = types.SvcBindingEqual
	BindingTypeFromString                = types.BindingTypeFromString
	RegisterCodec                        = types.RegisterCodec
	NewSvcDef                            = types.NewSvcDef
	MessagingTypeFromString              = types.MessagingTypeFromString
	OutputCachedEnumFromString           = types.OutputCachedEnumFromString
	OutputPrivacyEnumFromString          = types.OutputPrivacyEnumFromString
	ErrInvalidOutputPrivacyEnum          = types.ErrInvalidOutputPrivacyEnum
	ErrInvalidOutputCachedEnum           = types.ErrInvalidOutputCachedEnum
	ErrInvalidServiceName                = types.ErrInvalidServiceName
	ErrSvcBindingExists                  = types.ErrSvcBindingExists
	ErrInvalidPriceCount                 = types.ErrInvalidPriceCount
	ErrRefundDeposit                     = types.ErrRefundDeposit
	ErrLtMinProviderDeposit              = types.ErrLtMinProviderDeposit
	ErrLtServiceFee                      = types.ErrLtServiceFee
	ErrNotMatchingProvider               = types.ErrNotMatchingProvider
	ErrNotMatchingReqChainID             = types.ErrNotMatchingReqChainID
	NewGenesisState                      = types.NewGenesisState
	DefaultGenesisState                  = types.DefaultGenesisState
	ValidateGenesis                      = types.ValidateGenesis
	NewSvcRequest                        = types.NewSvcRequest
	ConvertRequestID                     = types.ConvertRequestID
	NewSvcResponse                       = types.NewSvcResponse
	NewReturnedFee                       = types.NewReturnedFee
	NewIncomingFee                       = types.NewIncomingFee
	GetServiceDefinitionKey              = types.GetServiceDefinitionKey
	GetMethodPropertyKey                 = types.GetMethodPropertyKey
	GetMethodsSubspaceKey                = types.GetMethodsSubspaceKey
	GetServiceBindingKey                 = types.GetServiceBindingKey
	GetBindingsSubspaceKey               = types.GetBindingsSubspaceKey
	GetRequestKey                        = types.GetRequestKey
	GetActiveRequestKey                  = types.GetActiveRequestKey
	GetSubActiveRequestKey               = types.GetSubActiveRequestKey
	GetResponseKey                       = types.GetResponseKey
	GetRequestsByExpirationIndexKeyByReq = types.GetRequestsByExpirationIndexKeyByReq
	GetRequestsByExpirationIndexKey      = types.GetRequestsByExpirationIndexKey
	GetRequestsByExpirationPrefix        = types.GetRequestsByExpirationPrefix
	GetReturnedFeeKey                    = types.GetReturnedFeeKey
	GetIncomingFeeKey                    = types.GetIncomingFeeKey
	PrometheusMetrics                    = types.PrometheusMetrics
	NopMetrics                           = types.NopMetrics
	NewMsgSvcDef                         = types.NewMsgSvcDef
	NewMsgSvcBind                        = types.NewMsgSvcBind
	NewMsgSvcBindingUpdate               = types.NewMsgSvcBindingUpdate
	NewMsgSvcDisable                     = types.NewMsgSvcDisable
	NewMsgSvcEnable                      = types.NewMsgSvcEnable
	NewMsgSvcRefundDeposit               = types.NewMsgSvcRefundDeposit
	NewMsgSvcRequest                     = types.NewMsgSvcRequest
	NewMsgSvcResponse                    = types.NewMsgSvcResponse
	NewMsgSvcRefundFees                  = types.NewMsgSvcRefundFees
	NewMsgSvcWithdrawFees                = types.NewMsgSvcWithdrawFees
	NewMsgSvcWithdrawTax                 = types.NewMsgSvcWithdrawTax
	NewParams                            = types.NewParams
	DefaultParams                        = types.DefaultParams
	MustUnmarshalParams                  = types.MustUnmarshalParams
	UnmarshalParams                      = types.UnmarshalParams
	MethodToMethodProperty               = types.MethodToMethodProperty

	// variable aliases
	ModuleCdc                    = types.ModuleCdc
	ServiceDefinitionKey         = types.ServiceDefinitionKey
	MethodPropertyKey            = types.MethodPropertyKey
	BindingPropertyKey           = types.BindingPropertyKey
	RequestKey                   = types.RequestKey
	ResponseKey                  = types.ResponseKey
	RequestsByExpirationIndexKey = types.RequestsByExpirationIndexKey
	IntraTxCounterKey            = types.IntraTxCounterKey
	ActiveRequestKey             = types.ActiveRequestKey
	ReturnedFeeKey               = types.ReturnedFeeKey
	IncomingFeeKey               = types.IncomingFeeKey
	ServiceFeeTaxKey             = types.ServiceFeeTaxKey
	ServiceSlashFractionKey      = types.ServiceSlashFractionKey
	DefaultMaxRequestTimeout     = types.DefaultMaxRequestTimeout
	DefaultMinDepositMultiple    = types.DefaultMinDepositMultiple
	DefaultServiceFeeTax         = types.DefaultServiceFeeTax
	DefaultSlashFraction         = types.DefaultSlashFraction
	DefaultComplaintRetrospect   = types.DefaultComplaintRetrospect
	DefaultArbitrationTimeLimit  = types.DefaultArbitrationTimeLimit
	DefaultTxSizeLimit           = types.DefaultTxSizeLimit
	MinRequestTimeout            = types.MinRequestTimeout
	MinDepositMultiple           = types.MinDepositMultiple
	MaxDepositMultiple           = types.MaxDepositMultiple
	MaxServiceFeeTax             = types.MaxServiceFeeTax
	MaxSlashFraction             = types.MaxSlashFraction
	MinComplaintRetrospect       = types.MinComplaintRetrospect
	MaxComplaintRetrospect       = types.MaxComplaintRetrospect
	MinArbitrationTimeLimit      = types.MinArbitrationTimeLimit
	MaxArbitrationTimeLimit      = types.MaxArbitrationTimeLimit
	MinTxSizeLimit               = types.MinTxSizeLimit
	MaxTxSizeLimit               = types.MaxTxSizeLimit
	KeyMaxRequestTimeout         = types.KeyMaxRequestTimeout
	KeyMinDepositMultiple        = types.KeyMinDepositMultiple
	KeyServiceFeeTax             = types.KeyServiceFeeTax
	KeySlashFraction             = types.KeySlashFraction
	KeyComplaintRetrospect       = types.KeyComplaintRetrospect
	KeyArbitrationTimeLimit      = types.KeyArbitrationTimeLimit
	KeyTxSizeLimit               = types.KeyTxSizeLimit
)

type (
	Keeper                = keeper.Keeper
	SvcBinding            = types.SvcBinding
	Level                 = types.Level
	BindingType           = types.BindingType
	SvcDef                = types.SvcDef
	MethodProperty        = types.MethodProperty
	OutputPrivacyEnum     = types.OutputPrivacyEnum
	OutputCachedEnum      = types.OutputCachedEnum
	MessagingType         = types.MessagingType
	GenesisState          = types.GenesisState
	SvcRequest            = types.SvcRequest
	SvcResponse           = types.SvcResponse
	ReturnedFee           = types.ReturnedFee
	IncomingFee           = types.IncomingFee
	Metrics               = types.Metrics
	MsgSvcDef             = types.MsgSvcDef
	MsgSvcBind            = types.MsgSvcBind
	MsgSvcBindingUpdate   = types.MsgSvcBindingUpdate
	MsgSvcDisable         = types.MsgSvcDisable
	MsgSvcEnable          = types.MsgSvcEnable
	MsgSvcRefundDeposit   = types.MsgSvcRefundDeposit
	MsgSvcRequest         = types.MsgSvcRequest
	MsgSvcResponse        = types.MsgSvcResponse
	MsgSvcRefundFees      = types.MsgSvcRefundFees
	MsgSvcWithdrawFees    = types.MsgSvcWithdrawFees
	MsgSvcWithdrawTax     = types.MsgSvcWithdrawTax
	Params                = types.Params
	QueryDefinitionParams = types.QueryDefinitionParams
	DefinitionOutput      = types.DefinitionOutput
	QueryBindingParams    = types.QueryBindingParams
	QueryBindingsParams   = types.QueryBindingsParams
	QueryRequestsParams   = types.QueryRequestsParams
	QueryResponseParams   = types.QueryResponseParams
	QueryFeesParams       = types.QueryFeesParams
	FeesOutput            = types.FeesOutput
)
