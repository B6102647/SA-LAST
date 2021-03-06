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
/**
 * 
 * @export
 * @interface ControllersBookEntity
 */
export interface ControllersBookEntity {
    /**
     * 
     * @type {number}
     * @memberof ControllersBookEntity
     */
    sid?: number;
}

export function ControllersBookEntityFromJSON(json: any): ControllersBookEntity {
    return ControllersBookEntityFromJSONTyped(json, false);
}

export function ControllersBookEntityFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersBookEntity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sid': !exists(json, 'sid') ? undefined : json['sid'],
    };
}

export function ControllersBookEntityToJSON(value?: ControllersBookEntity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sid': value.sid,
    };
}


