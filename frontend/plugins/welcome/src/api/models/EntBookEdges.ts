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
    EntBookBorrow,
    EntBookBorrowFromJSON,
    EntBookBorrowFromJSONTyped,
    EntBookBorrowToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBookEdges
 */
export interface EntBookEdges {
    /**
     * Booklist holds the value of the Booklist edge.
     * @type {Array<EntBookBorrow>}
     * @memberof EntBookEdges
     */
    booklist?: Array<EntBookBorrow>;
}

export function EntBookEdgesFromJSON(json: any): EntBookEdges {
    return EntBookEdgesFromJSONTyped(json, false);
}

export function EntBookEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBookEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'booklist': !exists(json, 'booklist') ? undefined : ((json['booklist'] as Array<any>).map(EntBookBorrowFromJSON)),
    };
}

export function EntBookEdgesToJSON(value?: EntBookEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'booklist': value.booklist === undefined ? undefined : ((value.booklist as Array<any>).map(EntBookBorrowToJSON)),
    };
}

