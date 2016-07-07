(function() {

    'use strict';

    class Attributes {
        constructor(spec) {
            this.family = spec.family;
            this.offense = spec.offense;
            this.defense = spec.defense;
            this.agility = spec.agility;
            this.range = spec.range;
            this.reproductivity = spec.reproductivity;
            this.energy = 1.0;
            this.hunger = 0.1;
        }
    }

    module.exports = Attributes;

}());
