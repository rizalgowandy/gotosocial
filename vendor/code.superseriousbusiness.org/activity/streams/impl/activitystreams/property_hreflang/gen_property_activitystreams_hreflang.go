// Code generated by astool. DO NOT EDIT.

package propertyhreflang

import (
	bcp47 "code.superseriousbusiness.org/activity/streams/values/bcp47"
	vocab "code.superseriousbusiness.org/activity/streams/vocab"
	"fmt"
	"net/url"
)

// ActivityStreamsHreflangProperty is the functional property "hreflang". It is
// permitted to be a single default-valued value type.
type ActivityStreamsHreflangProperty struct {
	rfcBcp47Member string
	hasBcp47Member bool
	unknown        interface{}
	iri            *url.URL
	alias          string
}

// DeserializeHreflangProperty creates a "hreflang" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeHreflangProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsHreflangProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "hreflang"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "hreflang")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &ActivityStreamsHreflangProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if v, err := bcp47.DeserializeBcp47(i); err == nil {
			this := &ActivityStreamsHreflangProperty{
				alias:          alias,
				hasBcp47Member: true,
				rfcBcp47Member: v,
			}
			return this, nil
		}
		this := &ActivityStreamsHreflangProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsHreflangProperty creates a new hreflang property.
func NewActivityStreamsHreflangProperty() *ActivityStreamsHreflangProperty {
	return &ActivityStreamsHreflangProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsRFCBcp47 afterwards
// will return false.
func (this *ActivityStreamsHreflangProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.hasBcp47Member = false
}

// Get returns the value of this property. When IsRFCBcp47 returns false, Get will
// return any arbitrary value.
func (this ActivityStreamsHreflangProperty) Get() string {
	return this.rfcBcp47Member
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this ActivityStreamsHreflangProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this ActivityStreamsHreflangProperty) HasAny() bool {
	return this.IsRFCBcp47() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this ActivityStreamsHreflangProperty) IsIRI() bool {
	return this.iri != nil
}

// IsRFCBcp47 returns true if this property is set and not an IRI.
func (this ActivityStreamsHreflangProperty) IsRFCBcp47() bool {
	return this.hasBcp47Member
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsHreflangProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsHreflangProperty) KindIndex() int {
	if this.IsRFCBcp47() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsHreflangProperty) LessThan(o vocab.ActivityStreamsHreflangProperty) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsRFCBcp47() && !o.IsRFCBcp47() {
		// Both are unknowns.
		return false
	} else if this.IsRFCBcp47() && !o.IsRFCBcp47() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsRFCBcp47() && o.IsRFCBcp47() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return bcp47.LessBcp47(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "hreflang".
func (this ActivityStreamsHreflangProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "hreflang"
	} else {
		return "hreflang"
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsHreflangProperty) Serialize() (interface{}, error) {
	if this.IsRFCBcp47() {
		return bcp47.SerializeBcp47(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsRFCBcp47 afterwards will return
// true.
func (this *ActivityStreamsHreflangProperty) Set(v string) {
	this.Clear()
	this.rfcBcp47Member = v
	this.hasBcp47Member = true
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *ActivityStreamsHreflangProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}
