// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getUsers(opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("proxmoxve:Permission/getUsers:getUsers", {
    }, opts);
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    readonly comments: string[];
    readonly emails: string[];
    readonly enableds: boolean[];
    readonly expirationDates: string[];
    readonly firstNames: string[];
    readonly groups: string[][];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keys: string[];
    readonly lastNames: string[];
    readonly userIds: string[];
}
