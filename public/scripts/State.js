(function() {

    'use strict';

    var glm = require('gl-matrix');

    class State {
        constructor(spec) {
            if (!spec) {
                throw 'No state argument';
            }
            this.type = spec.type;
            // physical
            this.size = spec.size;
            // energy
            this.energy = (spec.energy !== undefined) ? spec.energy : 1.0;
            // attacking
            this.target = spec.target;
            // position
            this.position = spec.position ? glm.vec3.fromValues(
                spec.position.x || spec.position[0] || 0,
                spec.position.y || spec.position[1] || 0,
                spec.position.z || spec.position[2] || 0) : glm.vec3.create();
            // rotation
            this.rotation = spec.rotation || 0;
        }
        interpolate(update, t) {
            var state = update.state;

            // iterpolate between current state and update based on a t value from 0 to 1
            // get distance vector
            let diff = glm.vec3.sub(glm.vec3.create(), state.position, this.position);
            // scale by t value
            diff = glm.vec3.scale(diff, diff, t);
            // get update position
            let position = glm.vec3.add(glm.vec3.create(), this.position, diff);

            // iterpolate size
            let size = this.size * (1-t) + state.size * t;

            return new State({
                type: state.type,
                position: position,
                size: size,
                rotation: this.rotation,
            });
        }
        update(update) {
            this.size = update.size;
            this.energy = update.energy;
            this.position = update.position;
            this.rotation = update.rotation;
            console.log(this);
        }
    }

    module.exports = State;

}());
