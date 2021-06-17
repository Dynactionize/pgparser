// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !cgo

// Package geographiclib is a wrapper around the GeographicLib library.
package geographiclib

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

var (
	// WGS84Spheroid represents the default WGS84 ellipsoid.
	WGS84Spheroid = NewSpheroid(6378137, 1/298.257223563)
)

// Spheroid is an object that can perform geodesic operations
// on a given spheroid.
type Spheroid struct {
	Radius       float64
	Flattening   float64
	SphereRadius float64
}

// NewSpheroid creates a spheroid from a radius and flattening.
func NewSpheroid(radius float64, flattening float64) *Spheroid {
	return &Spheroid{}
}

// Inverse solves the geodetic inverse problem on the given spheroid
// (https://en.wikipedia.org/wiki/Geodesy#Geodetic_problems).
// Returns s12 (distance in meters), az1 (azimuth at point 1) and az2 (azimuth at point 2).
func (s *Spheroid) Inverse(a, b s2.LatLng) (s12, az1, az2 float64) {
	return float64(0), float64(0), float64(0)
}

// InverseBatch computes the sum of the length of the lines represented
// by the line of points.
// This is intended for use for LineStrings. LinearRings/Polygons should use "AreaAndPerimeter".
// Returns the sum of the s12 (distance in meters) units.
func (s *Spheroid) InverseBatch(points []s2.Point) float64 {
	return float64(0)
}

// AreaAndPerimeter computes the area and perimeter of a polygon on a given spheroid.
// The points must never be duplicated (i.e. do not include the "final" point of a Polygon LinearRing).
// Area is in meter^2, Perimeter is in meters.
func (s *Spheroid) AreaAndPerimeter(points []s2.Point) (area float64, perimeter float64) {
	return float64(0), float64(0)
}

// Project returns computes the location of the projected point.
//
// Using the direct geodesic problem from GeographicLib (Karney 2013).
func (s *Spheroid) Project(point s2.LatLng, distance float64, azimuth s1.Angle) s2.LatLng {
	return s2.LatLngFromDegrees(float64(0), float64(0))
}
