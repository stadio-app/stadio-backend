//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	Address = Address.FromSchema(schema)
	AuthState = AuthState.FromSchema(schema)
	Country = Country.FromSchema(schema)
	EmailVerification = EmailVerification.FromSchema(schema)
	Event = Event.FromSchema(schema)
	Location = Location.FromSchema(schema)
	LocationSchedule = LocationSchedule.FromSchema(schema)
	Migration = Migration.FromSchema(schema)
	Owner = Owner.FromSchema(schema)
	Participant = Participant.FromSchema(schema)
	Review = Review.FromSchema(schema)
	SpatialRefSys = SpatialRefSys.FromSchema(schema)
	User = User.FromSchema(schema)
}
