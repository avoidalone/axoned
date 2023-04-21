package types

import (
	"fmt"
	"net/url"

	"cosmossdk.io/math"
)

// Parameter store keys.
var (
	ParamsKey      = []byte("Params")
	KeyInterpreter = []byte("Interpreter")
	KeyLimits      = []byte("Limits")
)

var (
	DefaultPredicatesWhitelist = make([]string, 0)
	DefaultPredicatesBlacklist = make([]string, 0)
	DefaultBootstrap           = ""
	DefaultMaxGas              = math.NewUint(uint64(100000))
	DefaultMaxSize             = math.NewUint(uint64(5000))
	DefaultMaxResultCount      = math.NewUint(uint64(1))
)

// NewParams creates a new Params object.
func NewParams(interpreter Interpreter, limits Limits) Params {
	return Params{
		Interpreter: interpreter,
		Limits:      limits,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(NewInterpreter(), DefaultLimits())
}

// Validate validates the set of params.
func (p Params) Validate() error {
	if err := validateInterpreter(p.Interpreter); err != nil {
		return err
	}
	if err := validateLimits(p.Limits); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	return p.Interpreter.String() + "\n" +
		p.Limits.String()
}

// NewInterpreter creates a new Interpreter with the given options.
func NewInterpreter(opts ...InterpreterOption) Interpreter {
	i := Interpreter{}
	for _, opt := range opts {
		opt(&i)
	}

	if i.PredicatesFilter.Whitelist == nil {
		i.PredicatesFilter.Whitelist = DefaultPredicatesWhitelist
	}

	if i.PredicatesFilter.Blacklist == nil {
		i.PredicatesFilter.Blacklist = DefaultPredicatesBlacklist
	}

	return i
}

// InterpreterOption is a functional option for configuring the Interpreter.
type InterpreterOption func(*Interpreter)

// WithPredicatesWhitelist sets the whitelist of predicates.
func WithPredicatesWhitelist(whitelist []string) InterpreterOption {
	return func(i *Interpreter) {
		i.PredicatesFilter.Whitelist = whitelist
	}
}

// WithPredicatesBlacklist sets the blacklist of predicates.
func WithPredicatesBlacklist(blacklist []string) InterpreterOption {
	return func(i *Interpreter) {
		i.PredicatesFilter.Blacklist = blacklist
	}
}

// WithBootstrap sets the bootstrap program.
func WithBootstrap(bootstrap string) InterpreterOption {
	return func(i *Interpreter) {
		i.Bootstrap = bootstrap
	}
}

func validateInterpreter(i interface{}) error {
	interpreter, ok := i.(Interpreter)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, file := range interpreter.VirtualFilesFilter.Whitelist {
		if _, err := url.Parse(file); err != nil {
			return fmt.Errorf("invalid virtual file in whitelist: %s", file)
		}
	}
	for _, file := range interpreter.VirtualFilesFilter.Blacklist {
		if _, err := url.Parse(file); err != nil {
			return fmt.Errorf("invalid virtual file in blacklist: %s", file)
		}
	}

	return nil
}

// NewLimits creates a new Limits object.
func NewLimits(maxGas, maxSize, maxResultCount *math.Uint) Limits {
	return Limits{
		MaxGas:         maxGas,
		MaxSize:        maxSize,
		MaxResultCount: maxResultCount,
	}
}

// DefaultLimits return a Limits object with default params.
func DefaultLimits() Limits {
	return NewLimits(&DefaultMaxGas, &DefaultMaxSize, &DefaultMaxResultCount)
}

func validateLimits(i interface{}) error {
	_, ok := i.(Limits)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// TODO: Validate limits params.
	return nil
}
