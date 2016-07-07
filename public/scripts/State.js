(function() {

    'use strict';

    var glm = require('gl-matrix');

    class State {
        constructor(spec) {
            this.type = spec.type;
            // health
            this.hunger = (spec.hunger !== undefined) ? spec.hunger : 0.0;
            this.energy = (spec.energy !== undefined) ? spec.energy : 1.0;
            // attacking / defending / consuming
            this.target = spec.target;
            // seeking / fleeing position
            this.position = spec.position ? glm.vec3.fromValues(
                spec.position.x || spec.position[0] || 0,
                spec.position.y || spec.position[1] || 0,
                spec.position.z || spec.position[2] || 0) : glm.vec3.create();
            this.rotation = spec.rotation || 0;
        }
        interpolate(update, t) {
            var state = update.state;

            // iterpolate between current state and update based on a t value from 0 to 1
            var position = this.position;

            switch (state.type) {
                case 'alive':
                    // get distance vector
                    let diff = glm.vec3.sub(glm.vec3.create(), state.position, this.position);
                    // scale by t value
                    diff = glm.vec3.scale(diff, diff, t);
                    // get update position
                    position = glm.vec3.add(glm.vec3.create(), this.position, diff);
                    break;
            }

            return new State({
                type: this.type,
                target: this.target,
                position: position,
                rotation: this.rotation
            });
        }
        update(update) {
            if (update.target) {
                this.target = update.target;
            }
            if (update.position) {
                this.position = update.position;
            }
            if (update.rotation) {
                this.rotation = update.rotation;
            }
        }
        color() {
            var health = (1 - this.hunger) * this.energy;
            switch (this.type) {
                case 'alive':
                    return [0.2 * health, 1.0 * health, 0.3 * health];
                case 'dead':
                    return [0.4, 0.4, 0.4];
                default:
                    return [1.0, 1.0, 0.0];
            }
        }
        matrix() {
            return glm.mat4.fromRotationTranslationScale(
                glm.mat4.create(),
                // rotation
                glm.quat.rotateZ(
                    glm.quat.create(),
                    glm.quat.identity(glm.quat.create()),
                    this.rotation),
                // translation
                this.position,
                // scale
                glm.vec3.fromValues(0.05, 0.05, 0.05));
        }
    }

    module.exports = State;

}());
