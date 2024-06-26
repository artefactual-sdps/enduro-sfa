/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { EnduroPackagePreservationAction } from './EnduroPackagePreservationAction';
import {
    EnduroPackagePreservationActionFromJSON,
    EnduroPackagePreservationActionFromJSONTyped,
    EnduroPackagePreservationActionToJSON,
} from './EnduroPackagePreservationAction';

/**
 * 
 * @export
 * @interface PreservationActionUpdatedEvent
 */
export interface PreservationActionUpdatedEvent {
    /**
     * Identifier of preservation action
     * @type {number}
     * @memberof PreservationActionUpdatedEvent
     */
    id: number;
    /**
     * 
     * @type {EnduroPackagePreservationAction}
     * @memberof PreservationActionUpdatedEvent
     */
    item: EnduroPackagePreservationAction;
}

/**
 * Check if a given object implements the PreservationActionUpdatedEvent interface.
 */
export function instanceOfPreservationActionUpdatedEvent(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "item" in value;

    return isInstance;
}

export function PreservationActionUpdatedEventFromJSON(json: any): PreservationActionUpdatedEvent {
    return PreservationActionUpdatedEventFromJSONTyped(json, false);
}

export function PreservationActionUpdatedEventFromJSONTyped(json: any, ignoreDiscriminator: boolean): PreservationActionUpdatedEvent {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'item': EnduroPackagePreservationActionFromJSON(json['item']),
    };
}

export function PreservationActionUpdatedEventToJSON(value?: PreservationActionUpdatedEvent | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'item': EnduroPackagePreservationActionToJSON(value.item),
    };
}

