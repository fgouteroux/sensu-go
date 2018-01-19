// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// TimeWindowWhenDaysFieldResolver implement to resolve requests for the TimeWindowWhen's days field.
type TimeWindowWhenDaysFieldResolver interface {
	// Days implements response to request for days field.
	Days(p graphql.ResolveParams) (interface{}, error)
}

//
// TimeWindowWhenFieldResolvers represents a collection of methods whose products represent the
// response values of the 'TimeWindowWhen' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type TimeWindowWhenFieldResolvers interface {
	TimeWindowWhenDaysFieldResolver
}

// TimeWindowWhenAliases implements all methods on TimeWindowWhenFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type TimeWindowWhenAliases struct{}

// Days implements response to request for 'days' field.
func (_ TimeWindowWhenAliases) Days(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// TimeWindowWhenType TimeWindowWhen defines the "when" attributes for time windows
var TimeWindowWhenType = graphql.NewType("TimeWindowWhen", graphql.ObjectKind)

// RegisterTimeWindowWhen registers TimeWindowWhen object type with given service.
func RegisterTimeWindowWhen(svc *graphql.Service, impl TimeWindowWhenFieldResolvers) {
	svc.RegisterObject(_ObjectTypeTimeWindowWhenDesc, impl)
}
func _ObjTypeTimeWindowWhenDaysHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowWhenDaysFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Days(p)
	}
}

func _ObjectTypeTimeWindowWhenConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "TimeWindowWhen defines the \"when\" attributes for time windows",
		Fields: graphql1.Fields{"days": &graphql1.Field{
			Args:              graphql1.FieldConfigArgument{},
			DeprecationReason: "",
			Description:       "Days is a hash of days",
			Name:              "days",
			Type:              graphql.OutputType("TimeWindowDays"),
		}},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see TimeWindowWhenFieldResolvers.")
		},
		Name: "TimeWindowWhen",
	}
}

// describe TimeWindowWhen's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeTimeWindowWhenDesc = graphql.ObjectDesc{
	Config:        _ObjectTypeTimeWindowWhenConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{"days": _ObjTypeTimeWindowWhenDaysHandler},
}

// TimeWindowDaysAllFieldResolver implement to resolve requests for the TimeWindowDays's all field.
type TimeWindowDaysAllFieldResolver interface {
	// All implements response to request for all field.
	All(p graphql.ResolveParams) (interface{}, error)
}

// TimeWindowDaysSundayFieldResolver implement to resolve requests for the TimeWindowDays's sunday field.
type TimeWindowDaysSundayFieldResolver interface {
	// Sunday implements response to request for sunday field.
	Sunday(p graphql.ResolveParams) (interface{}, error)
}

// TimeWindowDaysMondayFieldResolver implement to resolve requests for the TimeWindowDays's monday field.
type TimeWindowDaysMondayFieldResolver interface {
	// Monday implements response to request for monday field.
	Monday(p graphql.ResolveParams) (interface{}, error)
}

// TimeWindowDaysTuesdayFieldResolver implement to resolve requests for the TimeWindowDays's tuesday field.
type TimeWindowDaysTuesdayFieldResolver interface {
	// Tuesday implements response to request for tuesday field.
	Tuesday(p graphql.ResolveParams) (interface{}, error)
}

// TimeWindowDaysWednesdayFieldResolver implement to resolve requests for the TimeWindowDays's wednesday field.
type TimeWindowDaysWednesdayFieldResolver interface {
	// Wednesday implements response to request for wednesday field.
	Wednesday(p graphql.ResolveParams) (interface{}, error)
}

// TimeWindowDaysThursdayFieldResolver implement to resolve requests for the TimeWindowDays's thursday field.
type TimeWindowDaysThursdayFieldResolver interface {
	// Thursday implements response to request for thursday field.
	Thursday(p graphql.ResolveParams) (interface{}, error)
}

// TimeWindowDaysFridayFieldResolver implement to resolve requests for the TimeWindowDays's friday field.
type TimeWindowDaysFridayFieldResolver interface {
	// Friday implements response to request for friday field.
	Friday(p graphql.ResolveParams) (interface{}, error)
}

// TimeWindowDaysSaturdayFieldResolver implement to resolve requests for the TimeWindowDays's saturday field.
type TimeWindowDaysSaturdayFieldResolver interface {
	// Saturday implements response to request for saturday field.
	Saturday(p graphql.ResolveParams) (interface{}, error)
}

//
// TimeWindowDaysFieldResolvers represents a collection of methods whose products represent the
// response values of the 'TimeWindowDays' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type TimeWindowDaysFieldResolvers interface {
	TimeWindowDaysAllFieldResolver
	TimeWindowDaysSundayFieldResolver
	TimeWindowDaysMondayFieldResolver
	TimeWindowDaysTuesdayFieldResolver
	TimeWindowDaysWednesdayFieldResolver
	TimeWindowDaysThursdayFieldResolver
	TimeWindowDaysFridayFieldResolver
	TimeWindowDaysSaturdayFieldResolver
}

// TimeWindowDaysAliases implements all methods on TimeWindowDaysFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type TimeWindowDaysAliases struct{}

// All implements response to request for 'all' field.
func (_ TimeWindowDaysAliases) All(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// Sunday implements response to request for 'sunday' field.
func (_ TimeWindowDaysAliases) Sunday(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// Monday implements response to request for 'monday' field.
func (_ TimeWindowDaysAliases) Monday(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// Tuesday implements response to request for 'tuesday' field.
func (_ TimeWindowDaysAliases) Tuesday(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// Wednesday implements response to request for 'wednesday' field.
func (_ TimeWindowDaysAliases) Wednesday(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// Thursday implements response to request for 'thursday' field.
func (_ TimeWindowDaysAliases) Thursday(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// Friday implements response to request for 'friday' field.
func (_ TimeWindowDaysAliases) Friday(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// Saturday implements response to request for 'saturday' field.
func (_ TimeWindowDaysAliases) Saturday(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(interface{})
	return ret, err
}

// TimeWindowDaysType TimeWindowDays defines the days of a time window
var TimeWindowDaysType = graphql.NewType("TimeWindowDays", graphql.ObjectKind)

// RegisterTimeWindowDays registers TimeWindowDays object type with given service.
func RegisterTimeWindowDays(svc *graphql.Service, impl TimeWindowDaysFieldResolvers) {
	svc.RegisterObject(_ObjectTypeTimeWindowDaysDesc, impl)
}
func _ObjTypeTimeWindowDaysAllHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowDaysAllFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.All(p)
	}
}

func _ObjTypeTimeWindowDaysSundayHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowDaysSundayFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Sunday(p)
	}
}

func _ObjTypeTimeWindowDaysMondayHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowDaysMondayFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Monday(p)
	}
}

func _ObjTypeTimeWindowDaysTuesdayHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowDaysTuesdayFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Tuesday(p)
	}
}

func _ObjTypeTimeWindowDaysWednesdayHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowDaysWednesdayFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Wednesday(p)
	}
}

func _ObjTypeTimeWindowDaysThursdayHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowDaysThursdayFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Thursday(p)
	}
}

func _ObjTypeTimeWindowDaysFridayHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowDaysFridayFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Friday(p)
	}
}

func _ObjTypeTimeWindowDaysSaturdayHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowDaysSaturdayFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Saturday(p)
	}
}

func _ObjectTypeTimeWindowDaysConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "TimeWindowDays defines the days of a time window",
		Fields: graphql1.Fields{
			"all": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "all",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("TimeWindowTimeRange"))),
			},
			"friday": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "friday",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("TimeWindowTimeRange"))),
			},
			"monday": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "monday",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("TimeWindowTimeRange"))),
			},
			"saturday": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "saturday",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("TimeWindowTimeRange"))),
			},
			"sunday": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "sunday",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("TimeWindowTimeRange"))),
			},
			"thursday": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "thursday",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("TimeWindowTimeRange"))),
			},
			"tuesday": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "tuesday",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("TimeWindowTimeRange"))),
			},
			"wednesday": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "wednesday",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("TimeWindowTimeRange"))),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see TimeWindowDaysFieldResolvers.")
		},
		Name: "TimeWindowDays",
	}
}

// describe TimeWindowDays's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeTimeWindowDaysDesc = graphql.ObjectDesc{
	Config: _ObjectTypeTimeWindowDaysConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"all":       _ObjTypeTimeWindowDaysAllHandler,
		"friday":    _ObjTypeTimeWindowDaysFridayHandler,
		"monday":    _ObjTypeTimeWindowDaysMondayHandler,
		"saturday":  _ObjTypeTimeWindowDaysSaturdayHandler,
		"sunday":    _ObjTypeTimeWindowDaysSundayHandler,
		"thursday":  _ObjTypeTimeWindowDaysThursdayHandler,
		"tuesday":   _ObjTypeTimeWindowDaysTuesdayHandler,
		"wednesday": _ObjTypeTimeWindowDaysWednesdayHandler,
	},
}

// TimeWindowTimeRangeBeginFieldResolver implement to resolve requests for the TimeWindowTimeRange's begin field.
type TimeWindowTimeRangeBeginFieldResolver interface {
	// Begin implements response to request for begin field.
	Begin(p graphql.ResolveParams) (string, error)
}

// TimeWindowTimeRangeEndFieldResolver implement to resolve requests for the TimeWindowTimeRange's end field.
type TimeWindowTimeRangeEndFieldResolver interface {
	// End implements response to request for end field.
	End(p graphql.ResolveParams) (string, error)
}

//
// TimeWindowTimeRangeFieldResolvers represents a collection of methods whose products represent the
// response values of the 'TimeWindowTimeRange' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type TimeWindowTimeRangeFieldResolvers interface {
	TimeWindowTimeRangeBeginFieldResolver
	TimeWindowTimeRangeEndFieldResolver
}

// TimeWindowTimeRangeAliases implements all methods on TimeWindowTimeRangeFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type TimeWindowTimeRangeAliases struct{}

// Begin implements response to request for 'begin' field.
func (_ TimeWindowTimeRangeAliases) Begin(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(string)
	return ret, err
}

// End implements response to request for 'end' field.
func (_ TimeWindowTimeRangeAliases) End(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(string)
	return ret, err
}

// TimeWindowTimeRangeType TimeWindowTimeRange defines the time ranges of a time
var TimeWindowTimeRangeType = graphql.NewType("TimeWindowTimeRange", graphql.ObjectKind)

// RegisterTimeWindowTimeRange registers TimeWindowTimeRange object type with given service.
func RegisterTimeWindowTimeRange(svc *graphql.Service, impl TimeWindowTimeRangeFieldResolvers) {
	svc.RegisterObject(_ObjectTypeTimeWindowTimeRangeDesc, impl)
}
func _ObjTypeTimeWindowTimeRangeBeginHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowTimeRangeBeginFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Begin(p)
	}
}

func _ObjTypeTimeWindowTimeRangeEndHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(TimeWindowTimeRangeEndFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.End(p)
	}
}

func _ObjectTypeTimeWindowTimeRangeConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "TimeWindowTimeRange defines the time ranges of a time",
		Fields: graphql1.Fields{
			"begin": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Begin is the time which the time window should begin, in the format\n'3:00PM', which satisfies the time.Kitchen format",
				Name:              "begin",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"end": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "End is the time which the filter should end, in the format '3:00PM', which\nsatisfies the time.Kitchen format",
				Name:              "end",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see TimeWindowTimeRangeFieldResolvers.")
		},
		Name: "TimeWindowTimeRange",
	}
}

// describe TimeWindowTimeRange's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeTimeWindowTimeRangeDesc = graphql.ObjectDesc{
	Config: _ObjectTypeTimeWindowTimeRangeConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"begin": _ObjTypeTimeWindowTimeRangeBeginHandler,
		"end":   _ObjTypeTimeWindowTimeRangeEndHandler,
	},
}
