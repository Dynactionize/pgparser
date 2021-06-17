// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Package geos is a wrapper around the spatial data types between the geo
// package and the GEOS C library. The GEOS library is dynamically loaded
// at init time.
// Operations will error if the GEOS library was not found.

// +build !cgo

package geos

import (
	"github.com/ruiaylin/pgparser/types/geo/geopb"
)

// EnsureInitErrorDisplay is used to control the error message displayed by
// EnsureInit.
type EnsureInitErrorDisplay int

const (
	// EnsureInitErrorDisplayPrivate displays the full error message, including
	// path info. It is intended for log messages.
	EnsureInitErrorDisplayPrivate EnsureInitErrorDisplay = iota
	// EnsureInitErrorDisplayPublic displays a redacted error message, excluding
	// path info. It is intended for errors to display for the client.
	EnsureInitErrorDisplayPublic
)

// maxArrayLen is the maximum safe length for this architecture.
const maxArrayLen = 1<<31 - 1

// geosOnce contains the global instance of CR_GEOS, to be initialized
// during at a maximum of once.
// If it has failed to open, the error will be populated in "err".
// This should only be touched by "fetchGEOSOrError".
var geosOnce struct {
}

// EnsureInit attempts to start GEOS if it has not been opened already
// and returns the location if found, and an error if the CR_GEOS is not valid.
func EnsureInit(
	errDisplay EnsureInitErrorDisplay, flagLibraryDirectoryValue string,
) (string, error) {
	return "", nil
}

// A Error wraps an error returned from a GEOS operation.
type Error struct {
	msg string
}

// Error implements the error interface.
func (err *Error) Error() string {
	return err.msg
}

// WKTToEWKB parses a WKT into WKB using the GEOS library.
func WKTToEWKB(wkt geopb.WKT, srid geopb.SRID) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// BufferParamsJoinStyle maps to the GEOSBufJoinStyles enum in geos_c.h.in.
type BufferParamsJoinStyle int

// These should be kept in sync with the geos_c.h.in corresponding enum definition.
const (
	BufferParamsJoinStyleRound = 1
	BufferParamsJoinStyleMitre = 2
	BufferParamsJoinStyleBevel = 3
)

// BufferParamsEndCapStyle maps to the GEOSBufCapStyles enum in geos_c.h.in.
type BufferParamsEndCapStyle int

// These should be kept in sync with the geos_c.h.in corresponding enum definition.
const (
	BufferParamsEndCapStyleRound  = 1
	BufferParamsEndCapStyleFlat   = 2
	BufferParamsEndCapStyleSquare = 3
)

// BufferParams are parameters to provide into the GEOS buffer function.
type BufferParams struct {
	JoinStyle        BufferParamsJoinStyle
	EndCapStyle      BufferParamsEndCapStyle
	SingleSided      bool
	QuadrantSegments int
	MitreLimit       float64
}

// Buffer buffers the given geometry by the given distance and params.
func Buffer(ewkb geopb.EWKB, params BufferParams, distance float64) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// Area returns the area of an EWKB.
func Area(ewkb geopb.EWKB) (float64, error) {
	return float64(0), nil
}

// Length returns the length of an EWKB.
func Length(ewkb geopb.EWKB) (float64, error) {
	return float64(0), nil
}

// Centroid returns the centroid of an EWKB.
func Centroid(ewkb geopb.EWKB) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// ConvexHull returns an EWKB which returns the convex hull of the given EWKB.
func ConvexHull(ewkb geopb.EWKB) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// PointOnSurface returns an EWKB with a point that is on the surface of the given EWKB.
func PointOnSurface(ewkb geopb.EWKB) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// Intersection returns an EWKB which contains the geometries of intersection between A and B.
func Intersection(a geopb.EWKB, b geopb.EWKB) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// Union returns an EWKB which is a union of shapes A and B.
func Union(a geopb.EWKB, b geopb.EWKB) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// InterpolateLine returns the point along the given LineString which is at
// a given distance from starting point.
// Note: For distance less than 0 it returns start point similarly for distance
// greater LineString's length.
// InterpolateLine also works with (Multi)LineString. However, the result is
// not appropriate as it combines all the LineString present in (MULTI)LineString,
// considering all the corner points of LineString overlaps each other.
func InterpolateLine(ewkb geopb.EWKB, distance float64) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// MinDistance returns the minimum distance between two EWKBs.
func MinDistance(a geopb.EWKB, b geopb.EWKB) (float64, error) {
	return float64(0), nil
}

// ClipEWKBByRect clips a EWKB to the specified rectangle.
func ClipEWKBByRect(
	ewkb geopb.EWKB, xMin float64, yMin float64, xMax float64, yMax float64,
) (geopb.EWKB, error) {
	return geopb.EWKB{}, nil
}

// Covers returns whether the EWKB provided by A covers the EWKB provided by B.
func Covers(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

// CoveredBy returns whether the EWKB provided by A is covered by the EWKB provided by B.
func CoveredBy(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

// Contains returns whether the EWKB provided by A contains the EWKB provided by B.
func Contains(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

// Crosses returns whether the EWKB provided by A crosses the EWKB provided by B.
func Crosses(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

// Equals returns whether the EWKB provided by A equals the EWKB provided by B.
func Equals(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

// Intersects returns whether the EWKB provided by A intersects the EWKB provided by B.
func Intersects(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

// Overlaps returns whether the EWKB provided by A overlaps the EWKB provided by B.
func Overlaps(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

// Touches returns whether the EWKB provided by A touches the EWKB provided by B.
func Touches(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

// Within returns whether the EWKB provided by A is within the EWKB provided by B.
func Within(a geopb.EWKB, b geopb.EWKB) (bool, error) {
	return false, nil
}

//
// DE-9IM related
//

// Relate returns the DE-9IM relation between A and B.
func Relate(a geopb.EWKB, b geopb.EWKB) (string, error) {
	return "", nil
}

// RelatePattern whether A and B have a DE-9IM relation matching the given pattern.
func RelatePattern(a geopb.EWKB, b geopb.EWKB, pattern string) (bool, error) {
	return false, nil
}
