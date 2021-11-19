// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("proxmoxve:Permission/getUser:getUser", {
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    userId: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    readonly acls: outputs.Permission.GetUserAcl[];
    readonly comment: string;
    readonly email: string;
    readonly enabled: boolean;
    readonly expirationDate: string;
    readonly firstName: string;
    readonly groups: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keys: string;
    readonly lastName: string;
    readonly userId: string;
}

export function getUserOutput(args: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    return pulumi.output(args).apply(a => getUser(a, opts))
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    userId: pulumi.Input<string>;
}
