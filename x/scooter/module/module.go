package module

import (
	"context"
	"encoding/json"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/semalis/uram/x/scooter/keeper"
	"github.com/semalis/uram/x/scooter/types"
)

var (
	_ appmodule.AppModule = AppModule{}
)

type AppModule struct {
	cdc    codec.Codec
	keeper keeper.Keeper
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

func (am AppModule) InitGenesis(ctx context.Context, cdc codec.JSONCodec, data json.RawMessage) []module.ValidatorUpdate {
	return nil
}

func (am AppModule) ExportGenesis(ctx context.Context, cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

func (am AppModule) AppModuleBasic() module.AppModuleBasic {
	return AppModuleBasic{}
}

func (AppModule) IsOnePerModuleType() {}

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string { return types.ModuleName }

func (AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	return cdc.UnmarshalJSON(bz, &data)
}

// ProvideModule for depinject
func ProvideModule(cdc codec.Codec, k keeper.Keeper) appmodule.AppModule {
	return AppModule{cdc: cdc, keeper: k}
}
