/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Package not found
 * @export
 * @interface PackageConfirmNotFoundResponseBody
 */
export interface PackageConfirmNotFoundResponseBody {
    /**
     * Identifier of missing package
     * @type {number}
     * @memberof PackageConfirmNotFoundResponseBody
     */
    id: number;
    /**
     * Message of error
     * @type {string}
     * @memberof PackageConfirmNotFoundResponseBody
     */
    message: string;
}

/**
 * Check if a given object implements the PackageConfirmNotFoundResponseBody interface.
 */
export function instanceOfPackageConfirmNotFoundResponseBody(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "message" in value;

    return isInstance;
}

export function PackageConfirmNotFoundResponseBodyFromJSON(json: any): PackageConfirmNotFoundResponseBody {
    return PackageConfirmNotFoundResponseBodyFromJSONTyped(json, false);
}

export function PackageConfirmNotFoundResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackageConfirmNotFoundResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'message': json['message'],
    };
}

export function PackageConfirmNotFoundResponseBodyToJSON(value?: PackageConfirmNotFoundResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'message': value.message,
    };
}

