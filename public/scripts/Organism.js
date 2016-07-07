(function() {

    'use strict';

    class Organism {
        constructor(spec) {
            this.id = spec.id;
            this.position = spec.position;
            this.rotation = spec.rotation;
            this.state = spec.state;
            this.attributes = spec.attributes;
        }
        interpolate(update, t) {
            // iterpolate between current state and update based on a t value from 0 to 1
            console.log(update, t);
        }
        apply(update) {
            // apply an update to the organism
            console.log(update);
        }
    }

    module.exports = Organism;

}());
