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
            this.agility = spec.agility;
            this.reproductivity = spec.reproductivity;
            this.growthRate = spec.growthRate;
            // coordniate based
            this.offspringsize = spec.offspringsize;
            this.range = spec.range;
            this.perception = spec.perception;
            this.speed = spec.speed;
        }
    }

    module.exports = Attributes;

}());
