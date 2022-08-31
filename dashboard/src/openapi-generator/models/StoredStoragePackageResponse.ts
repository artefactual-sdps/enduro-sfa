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
 * A StoredStoragePackage describes a package retrieved by the storage service. (default view)
 * @export
 * @interface StoredStoragePackageResponse
 */
export interface StoredStoragePackageResponse {
    /**
     * 
     * @type {string}
     * @memberof StoredStoragePackageResponse
     */
    aipId: string;
    /**
     * Creation datetime
     * @type {Date}
     * @memberof StoredStoragePackageResponse
     */
    createdAt: Date;
    /**
     * 
     * @type {string}
     * @memberof StoredStoragePackageResponse
     */
    locationId?: string;
    /**
     * 
     * @type {string}
     * @memberof StoredStoragePackageResponse
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof StoredStoragePackageResponse
     */
    objectKey: string;
    /**
     * Status of the package
     * @type {string}
     * @memberof StoredStoragePackageResponse
     */
    status: StoredStoragePackageResponseStatusEnum;
}


/**
 * @export
 */
export const StoredStoragePackageResponseStatusEnum = {
    Unspecified: 'unspecified',
    InReview: 'in_review',
    Rejected: 'rejected',
    Stored: 'stored',
    Moving: 'moving'
} as const;
export type StoredStoragePackageResponseStatusEnum = typeof StoredStoragePackageResponseStatusEnum[keyof typeof StoredStoragePackageResponseStatusEnum];


export function StoredStoragePackageResponseFromJSON(json: any): StoredStoragePackageResponse {
    return StoredStoragePackageResponseFromJSONTyped(json, false);
}

export function StoredStoragePackageResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): StoredStoragePackageResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'aipId': json['aip_id'],
        'createdAt': (new Date(json['created_at'])),
        'locationId': !exists(json, 'location_id') ? undefined : json['location_id'],
        'name': json['name'],
        'objectKey': json['object_key'],
        'status': json['status'],
    };
}

export function StoredStoragePackageResponseToJSON(value?: StoredStoragePackageResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'aip_id': value.aipId,
        'created_at': (value.createdAt.toISOString()),
        'location_id': value.locationId,
        'name': value.name,
        'object_key': value.objectKey,
        'status': value.status,
    };
}
