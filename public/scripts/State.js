(function() {

    'use strict';

    class State {
        constructor(spec) {
            this.type = spec.type;
            // attacking / defending / consuming
            this.target = spec.target;
            // seeking / fleeing position
            this.x = spec.x;
            this.y = spec.y;
        }
    }

    module.exports = State;

}());
