package model

// Package model has types used to build a tutorial from discovered markdown.
//
// The file hierarchy holding the markdown defines tutorial structure,
// organizing markdown files (lessons) into nestable groups (courses).
//
// Suppose one offers a tutorial on Benelux - an economic union of
// three neighbouring states in western Europe.
//
// The first lesson is an overview of Benelux, with sibling (_not_ child)
// courses covering Belgium, Netherlands, and Luxembourg (as one might arrange
// a physical book).  These courses may hold lessons on provinces, or sub-courses
// drilling into cities, regional histories etc.  This hierarchy is optional, but
// supported and helpful to organize material.
//
// Associated content REST addresses reflect the file system hierarchy, e.g.
//
//     benelux.com/overview                  // Describes Benelux in general.
//     benelux.com/history                   // Benelux history, economy, etc.
//     benelux.com/economy
//     benelux.com/belgium/overview          // Describes Belgium in general.
//     benelux.com/belgium/tintin            // Dive into important details.
//     benelux.com/belgium/beer
//     benelux.com/belgium/antwerp/overview  // Dive into Antwerp, etc.
//     benelux.com/belgium/antwerp/diamonds
//     benelux.com/belgium/antwerp/rubens
//     benelux.com/belgium/east-flanders
//     benelux.com/belgium/brabant
//     ...
//     benelux.com/netherlands/overview
//     benelux.com/netherlands/drenthe
//     benelux.com/netherlands/flevoland
//     ...
//
// All content is accessible from standard nav tools:
//
//      overview     |                           {main page outline
//      belgium      |                            here - title, h1,
//     [netherlands] |       {main page           h2, h3 etc.}
//      luxembourg   |        content here}
//
// At all times exactly one of the left nav choices is selected., and the
// main page shows content associated with that selection.
//
// The first item, in this case "overview", is the initial highlight.
// If one hits the domain without a REST path, one is redirected to
// /overview and that item is highlighted in the menu, and its
// content is shown.
//
// Items in the left nav either name content and show it when clicked, or
// they name sub-courses and expand choices when clicked.
// In the latter case, the main content and the left nav highlighting
// don't change.  A second click hides the exposed sub-course names.
//
// Only the name of a lesson (a leaf) with content can 1) be highlighted,
// 2) change the main page content when clicked, and 3) serve at a meaningful
// REST address.  Everything else is a course, and only expands or hides
// its own appearance.
//
// By design, this scheme maps to this filesystem layout:
//
//     benelux/
//       01_history.md
//       02_economy.md
//       README.md
//       03_belgium/
//         01_tintin.md
//         02_beer.md
//         03_antwerp/
//           README.md
//           01_diamonds.md
//           ...
//         04_east-flanders.md
//         05_brabant.md
//         ...
//       04_netherlands/
//         README.md
//         01_drenthe.md
//         02_flevoland.md
//       ...
//
// Where, say README (a github name convention) is converted to "overview"
// by a file loader, and likewise leading numbers in file names are dropped
// - though the implied presentation order is preserved in the nav so one
// can retain a lesson ordering.
//
// The proposed command line to read and serve content is
//
//      mdrip --mode web /foo/benelux
// or
//      mdrip --mode web /foo/benelux/README.md
//
// The arg names either a directory or a file.
//
// If the arg is a directory name, the tree below it is read in an attempt
// to build REST-fully addressable content and UX.  The names shown in the UX
// could be raw file names or could be processed a bit, e.g. underscores or
// hyphens become spaces, the ordering of the content in the UX could be
// controlled by omittable numerical prefixes on file names, etc.
// Errors in tree structure dealt with reasonably or cause immediate server
// failure.
//
// If only one file is read, then only that content is shown -
// no left nav needed.
