var gulp = require('gulp');
var browserify = require('gulp-browserify');
var less = require('gulp-less');
var path = require('path');

gulp.task('js', function() {
  gulp.src('src/js/main.js')
    .pipe(browserify({
      insertGlobals : true,
      basedir: "."
    }))
    .pipe(
      gulp.dest('./public/js')
    );
});

gulp.task('less', function () {
  gulp.src('./src/less/main.less')
    .pipe(less({
      paths: [ path.join(__dirname, 'src/less', 'src/less/includes') ]
    }))
    .pipe(gulp.dest('./public/css'));
});

gulp.task('watch', function() {
  gulp.watch('src/js/*', ['js']);
  gulp.watch('src/less/*', ['less']);
});

gulp.task('default', ['js', 'less']);
