import _ from 'lodash'
export function clearObj(params) {
    _.mapKeys(params, function (value, key) {
        // console.log(`${value} - ${key}`)
        if (_.isEmpty(value)) delete params[key]
    })
    return params;
}
