(function() {

	'use strict';

	var gulp = require('gulp');
	var concat = require('gulp-concat');
	var source = require('vinyl-source-stream');
	var watchify = require('watchify');
	var del = require('del');
	var jshint = require('gulp-jshint');
	var csso = require('gulp-csso');
	var runSequence = require('run-sequence');
	var gulpgo = require('gulp-go');
	var browserify = require('browserify');
	var babel = require('babelify');

	var PROJECT_NAME = 'egol';
	var PUBLIC_DIR = './public/';
	var OUTPUT_DIR = './build/public/';
	var paths = {
		serverRoot: './main.go',
		webappRoot: PUBLIC_DIR + '/app.js',
		go: [ './api/**/*.go', './main.go' ],
		templates: [ PUBLIC_DIR + 'templates/**/*.hbs'],
		scripts: [ PUBLIC_DIR + 'scripts/**/*.js',  PUBLIC_DIR + 'app.js' ],
		styles: [ PUBLIC_DIR + 'styles/reset.css',  PUBLIC_DIR + 'styles/**/*.css' ],
		index: [ PUBLIC_DIR + 'index.html' ],
		resources: [
			PUBLIC_DIR + 'index.html',
			PUBLIC_DIR + 'shaders/*',
			PUBLIC_DIR + 'favicons/*',
			PUBLIC_DIR + 'images/*'
		]
	};

	function handleError(err) {
		console.log(err);
		this.emit('end');
	}

	function bundle(bundler, watch) {
		if (watch) {
			var watcher = watchify(bundler);
			watcher.on('update', function(ids) {
				// When any files updates
				console.log('\nWatch detected changes to: ');
				for (var i=0; i<ids.length; ids++) {
				   console.log('\t'+ids[i]);
				}
				var updateStart = Date.now();
				watcher.bundle()
					.on('error', handleError)
					.pipe(source(PROJECT_NAME + '.js'))
					// This is where you add uglifying etc.
					.pipe(gulp.dest(OUTPUT_DIR));
				console.log('Updated in', (Date.now() - updateStart) + 'ms');
			});
			bundler = watcher;
		}
		return bundler
			.bundle() // Create the initial bundle when starting the task
			.on('error', handleError)
			.pipe(source(PROJECT_NAME + '.js'))
			.pipe(gulp.dest(OUTPUT_DIR));
	}

	gulp.task('clean', function(done) {
		del.sync([ OUTPUT_DIR ]);
		done();
	});

	gulp.task('lint', function() {
		return gulp.src([ PUBLIC_DIR + '/**/*.js',
			'!'+PUBLIC_DIR+'/extern/**/*.js'])
			.pipe(jshint('.jshintrc'))
			.pipe(jshint.reporter('jshint-stylish'));
	});

	gulp.task('build-and-watch-scripts', function() {
		var browserify = require('browserify'),
			bundler = browserify(paths.webappRoot, {
				debug: true,
				standalone: PROJECT_NAME
			}).transform(babel, {presets: ['es2015']});
		return bundle(bundler, true);
	});

	gulp.task('build-scripts', function() {
		var bundler = browserify(paths.webappRoot, {
			debug: true,
			standalone: PROJECT_NAME
		});
		return bundle(bundler, false);
	});

	gulp.task('build-styles', function () {
		return gulp.src(paths.styles)
			.pipe(csso())
			.pipe(concat(PROJECT_NAME + '.css'))
			.pipe(gulp.dest(OUTPUT_DIR));
	});

	gulp.task('copy-resources', function() {
		return gulp.src(paths.resources, {
				base: PUBLIC_DIR
			})
			.pipe(gulp.dest(OUTPUT_DIR));
	});

	gulp.task('build-watch', function(done) {
		runSequence(
			[ 'clean', 'lint' ],
			[ 'build-and-watch-scripts', 'build-styles', 'copy-resources' ],
			done);
	});

	gulp.task('build', function(done) {
		runSequence(
			[ 'clean', 'lint' ],
			[ 'build-scripts', 'build-styles', 'copy-resources' ],
			done);
	});

	var go;
	gulp.task('serve', function() {
		go = gulpgo.run(paths.serverRoot, [], {
			cwd: __dirname,
			stdio: 'inherit'
		});
	});

	gulp.task('watch', [ 'build-watch' ], function(done) {
		gulp.watch(paths.go).on('change', function() {
			go.restart();
		});
		gulp.watch(paths.styles, [ 'build-styles' ]);
		gulp.watch(paths.resources, [ 'copy-resources' ]);
		done();
	});

	gulp.task('default', [ 'watch', 'serve' ], function() {
	});

}());
