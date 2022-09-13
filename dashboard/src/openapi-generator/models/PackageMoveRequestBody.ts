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
 * 
 * @export
 * @interface PackageMoveRequestBody
 */
export interface PackageMoveRequestBody {
    /**
     * 
     * @type {string}
     * @memberof PackageMoveRequestBody
     */
    locationId: string;
}

/**
 * Check if a given object implements the PackageMoveRequestBody interface.
 */
export function instanceOfPackageMoveRequestBody(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "locationId" in value;

    return isInstance;
}

export function PackageMoveRequestBodyFromJSON(json: any): PackageMoveRequestBody {
    return PackageMoveRequestBodyFromJSONTyped(json, false);
}

export function PackageMoveRequestBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackageMoveRequestBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'locationId': json['location_id'],
    };
}

export function PackageMoveRequestBodyToJSON(value?: PackageMoveRequestBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'location_id': value.locationId,
    };
}

