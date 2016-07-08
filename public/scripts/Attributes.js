(function() {

    'use strict';

    class Attributes {
        constructor(spec) {
            if (!spec) {
                throw 'No attribute argument';
            }
            this.family = spec.family;
            this.offense = spec.offense;
            this.defense = spec.defense;
            this.reproductivity = spec.reproductivity;
            // coordniate based
            this.range = spec.range;
            this.perception = spec.perception;
            this.speed = spec.speed;
        }
    }

    module.exports = Attributes;

}());
