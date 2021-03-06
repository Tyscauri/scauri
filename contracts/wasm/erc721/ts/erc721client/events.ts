// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmclient from "wasmclient"

const erc721Handlers = new Map<string, (evt: Erc721Events, msg: string[]) => void>([
	["erc721.approval", (evt: Erc721Events, msg: string[]) => evt.approval(new EventApproval(msg))],
	["erc721.approvalForAll", (evt: Erc721Events, msg: string[]) => evt.approvalForAll(new EventApprovalForAll(msg))],
	["erc721.init", (evt: Erc721Events, msg: string[]) => evt.init(new EventInit(msg))],
	["erc721.mint", (evt: Erc721Events, msg: string[]) => evt.mint(new EventMint(msg))],
	["erc721.transfer", (evt: Erc721Events, msg: string[]) => evt.transfer(new EventTransfer(msg))],
]);

export class Erc721Events implements wasmclient.IEventHandler {
/* eslint-disable @typescript-eslint/no-empty-function */
	approval: (evt: EventApproval) => void = () => {};
	approvalForAll: (evt: EventApprovalForAll) => void = () => {};
	init: (evt: EventInit) => void = () => {};
	mint: (evt: EventMint) => void = () => {};
	transfer: (evt: EventTransfer) => void = () => {};
/* eslint-enable @typescript-eslint/no-empty-function */

	public callHandler(topic: string, params: string[]): void {
		const handler = erc721Handlers.get(topic);
		if (handler) {
			handler(this, params);
		}
	}

	public onErc721Approval(handler: (evt: EventApproval) => void): void {
		this.approval = handler;
	}

	public onErc721ApprovalForAll(handler: (evt: EventApprovalForAll) => void): void {
		this.approvalForAll = handler;
	}

	public onErc721Init(handler: (evt: EventInit) => void): void {
		this.init = handler;
	}

	public onErc721Mint(handler: (evt: EventMint) => void): void {
		this.mint = handler;
	}

	public onErc721Transfer(handler: (evt: EventTransfer) => void): void {
		this.transfer = handler;
	}
}

export class EventApproval extends wasmclient.Event {
	public readonly approved: wasmclient.AgentID;
	public readonly owner: wasmclient.AgentID;
	public readonly tokenID: wasmclient.Hash;
	
	public constructor(msg: string[]) {
		super(msg);
		this.approved = this.nextAgentID();
		this.owner = this.nextAgentID();
		this.tokenID = this.nextHash();
	}
}

export class EventApprovalForAll extends wasmclient.Event {
	public readonly approval: wasmclient.Bool;
	public readonly operator: wasmclient.AgentID;
	public readonly owner: wasmclient.AgentID;
	
	public constructor(msg: string[]) {
		super(msg);
		this.approval = this.nextBool();
		this.operator = this.nextAgentID();
		this.owner = this.nextAgentID();
	}
}

export class EventInit extends wasmclient.Event {
	public readonly name: wasmclient.String;
	public readonly symbol: wasmclient.String;
	
	public constructor(msg: string[]) {
		super(msg);
		this.name = this.nextString();
		this.symbol = this.nextString();
	}
}

export class EventMint extends wasmclient.Event {
	public readonly balance: wasmclient.Uint64;
	public readonly owner: wasmclient.AgentID;
	public readonly tokenID: wasmclient.Hash;
	
	public constructor(msg: string[]) {
		super(msg);
		this.balance = this.nextUint64();
		this.owner = this.nextAgentID();
		this.tokenID = this.nextHash();
	}
}

export class EventTransfer extends wasmclient.Event {
	public readonly from: wasmclient.AgentID;
	public readonly to: wasmclient.AgentID;
	public readonly tokenID: wasmclient.Hash;
	
	public constructor(msg: string[]) {
		super(msg);
		this.from = this.nextAgentID();
		this.to = this.nextAgentID();
		this.tokenID = this.nextHash();
	}
}
