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
import {
    EnduroStoredPackageResponseBody,
    EnduroStoredPackageResponseBodyFromJSON,
    EnduroStoredPackageResponseBodyFromJSONTyped,
    EnduroStoredPackageResponseBodyToJSON,
} from './EnduroStoredPackageResponseBody';

/**
 * EnduroPackage-Updated-EventResponseBody result type (default view)
 * @export
 * @interface EnduroPackageUpdatedEventResponseBody
 */
export interface EnduroPackageUpdatedEventResponseBody {
    /**
     * Identifier of package
     * @type {number}
     * @memberof EnduroPackageUpdatedEventResponseBody
     */
    id: number;
    /**
     * 
     * @type {EnduroStoredPackageResponseBody}
     * @memberof EnduroPackageUpdatedEventResponseBody
     */
    item: EnduroStoredPackageResponseBody;
}

export function EnduroPackageUpdatedEventResponseBodyFromJSON(json: any): EnduroPackageUpdatedEventResponseBody {
    return EnduroPackageUpdatedEventResponseBodyFromJSONTyped(json, false);
}

export function EnduroPackageUpdatedEventResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroPackageUpdatedEventResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'item': EnduroStoredPackageResponseBodyFromJSON(json['item']),
    };
}

export function EnduroPackageUpdatedEventResponseBodyToJSON(value?: EnduroPackageUpdatedEventResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'item': EnduroStoredPackageResponseBodyToJSON(value.item),
    };
}
