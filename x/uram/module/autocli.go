package uram

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"uram/x/uram/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListScooter",
					Use:       "list-scooter",
					Short:     "List all scooter",
				},
				{
					RpcMethod:      "GetScooter",
					Use:            "get-scooter [id]",
					Short:          "Gets a scooter by id",
					Alias:          []string{"show-scooter"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateScooter",
					Use:            "create-scooter [location] [active] [owner] [price]",
					Short:          "Create a new scooter",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "location"}, {ProtoField: "active"}, {ProtoField: "owner"}, {ProtoField: "price"}},
				},
				{
					RpcMethod:      "UpdateScooter",
					Use:            "update-scooter [id] [location] [active] [owner] [price]",
					Short:          "Update an existing scooter",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "location"}, {ProtoField: "active"}, {ProtoField: "owner"}, {ProtoField: "price"}},
				},
				{
					RpcMethod:      "DeleteScooter",
					Use:            "delete-scooter [id]",
					Short:          "Delete scooter",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
