/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntRoleEdges,
    EntRoleEdgesFromJSON,
    EntRoleEdgesFromJSONTyped,
    EntRoleEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRole
 */
export interface EntRole {
    /**
     * ROLENAME holds the value of the "ROLE_NAME" field.
     * @type {string}
     * @memberof EntRole
     */
    rOLENAME?: string;
    /**
     * 
     * @type {EntRoleEdges}
     * @memberof EntRole
     */
    edges?: EntRoleEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntRole
     */
    id?: number;
}

export function EntRoleFromJSON(json: any): EntRole {
    return EntRoleFromJSONTyped(json, false);
}

export function EntRoleFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRole {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'rOLENAME': !exists(json, 'ROLE_NAME') ? undefined : json['ROLE_NAME'],
        'edges': !exists(json, 'edges') ? undefined : EntRoleEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntRoleToJSON(value?: EntRole | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ROLE_NAME': value.rOLENAME,
        'edges': EntRoleEdgesToJSON(value.edges),
        'id': value.id,
    };
}


