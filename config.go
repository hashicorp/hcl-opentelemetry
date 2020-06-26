package hclotel

import (
	"time"
)

const (
	// DefaultPeriod is the default period to emit metrics.
	DefaultPeriod time.Duration = time.Minute

	// DefaultPrometheusPort is the default port for HTTP service to bind on.
	DefaultPrometheusPort uint = 8888
)

// TelemetryConfig is the configuration for telemetry.
type TelemetryConfig struct {
	MetricsPrefix *string `mapstructure:"metrics_prefix"`

	Stdout     *StdoutConfig     `mapstructure:"stdout"`
	DogStatsD  *DogStatsDConfig  `mapstructure:"dogstatsd"`
	Prometheus *PrometheusConfig `mapstructure:"prometheus"`
}

// StdoutConfig is the configuration for emitting metrics to stdout.
type StdoutConfig struct {
	Period         *time.Duration `mapstructure:"period"`
	PrettyPrint    *bool          `mapstructure:"pretty_print"`
	DoNotPrintTime *bool          `mapstructure:"do_not_print_time"`
}

// DogStatsDConfig is the configuration for emitting metrics to dogstatsd.
type DogStatsDConfig struct {
	Address *string        `mapstructure:"address"`
	Period  *time.Duration `mapstructure:"period"`
}

// PrometheusConfig is the configuration for emitting metrics to Prometheus.
type PrometheusConfig struct {
	Port        *uint          `mapstructure:"port"`
	CachePeriod *time.Duration `mapstructure:"cache_period"`
}

// DefaultTelemetryConfig returns a configuration that is populated with the
// default values.
func DefaultTelemetryConfig() *TelemetryConfig {
	return &TelemetryConfig{}
}

// Copy returns a deep copy of this configuration.
func (c *TelemetryConfig) Copy() *TelemetryConfig {
	if c == nil {
		return nil
	}

	var o TelemetryConfig
	o.MetricsPrefix = StringCopy(c.MetricsPrefix)

	if c.Stdout != nil {
		o.Stdout = c.Stdout.Copy()
	}
	if c.DogStatsD != nil {
		o.DogStatsD = c.DogStatsD.Copy()
	}
	if c.Prometheus != nil {
		o.Prometheus = c.Prometheus.Copy()
	}

	return &o
}

// Merge combines all values in this configuration with the values in the other
// configuration, with values in the other configuration taking precedence.
// Maps and slices are merged, most other values are overwritten. Complex
// structs define their own merge functionality
func (c *TelemetryConfig) Merge(o *TelemetryConfig) *TelemetryConfig {
	if c == nil {
		if o == nil {
			return nil
		}
		return o.Copy()
	}

	if o == nil {
		return c.Copy()
	}

	r := c.Copy()

	if o.MetricsPrefix != nil {
		r.MetricsPrefix = StringCopy(o.MetricsPrefix)
	}

	if o.Stdout != nil {
		r.Stdout = o.Stdout.Copy()
	}

	if o.DogStatsD != nil {
		r.DogStatsD = o.DogStatsD.Copy()
	}

	if o.Prometheus != nil {
		r.Prometheus = o.Prometheus.Copy()
	}

	return r
}

// Finalize ensures there no nested nil pointers.
func (c *TelemetryConfig) Finalize() {
	if c == nil {
		return
	}

	d := DefaultTelemetryConfig()

	if c.MetricsPrefix == nil {
		c.MetricsPrefix = d.MetricsPrefix
	}

	c.Stdout.Finalize()
	c.DogStatsD.Finalize()
	c.Prometheus.Finalize()
}

// DefaultStdoutConfig returns a configuration that is populated with the
// default values.
func DefaultStdoutConfig() *StdoutConfig {
	return &StdoutConfig{
		Period:         TimeDuration(DefaultPeriod),
		PrettyPrint:    Bool(false),
		DoNotPrintTime: Bool(false),
	}
}

// Copy returns a deep copy of this configuration.
func (c *StdoutConfig) Copy() *StdoutConfig {
	if c == nil {
		return nil
	}

	return &StdoutConfig{
		Period:         TimeDurationCopy(c.Period),
		PrettyPrint:    BoolCopy(c.PrettyPrint),
		DoNotPrintTime: BoolCopy(c.DoNotPrintTime),
	}
}

// Finalize ensures there no nil pointers.
func (c *StdoutConfig) Finalize() {
	if c == nil {
		return
	}

	d := DefaultStdoutConfig()

	if c.Period == nil {
		c.Period = d.Period
	}

	if c.PrettyPrint == nil {
		c.PrettyPrint = d.PrettyPrint
	}

	if c.DoNotPrintTime == nil {
		c.DoNotPrintTime = d.DoNotPrintTime
	}
}

// DefaultDogStatsDConfig returns a configuration that is populated with the
// default values.
func DefaultDogStatsDConfig() *DogStatsDConfig {
	return &DogStatsDConfig{
		Address: String("udp://127.0.0.1:8125"),
		Period:  TimeDuration(DefaultPeriod),
	}
}

// Merge combines all values in this configuration with the values in the other
// configuration, with values in the other configuration taking precedence.
// Maps and slices are merged, most other values are overwritten. Complex
// structs define their own merge functionality.
func (c *StdoutConfig) Merge(o *StdoutConfig) *StdoutConfig {
	if c == nil {
		if o == nil {
			return nil
		}
		return o.Copy()
	}

	if o == nil {
		return c.Copy()
	}

	r := c.Copy()

	if o.Period != nil {
		r.Period = TimeDurationCopy(o.Period)
	}

	if o.PrettyPrint != nil {
		r.PrettyPrint = BoolCopy(o.PrettyPrint)
	}

	if o.DoNotPrintTime != nil {
		r.DoNotPrintTime = BoolCopy(o.DoNotPrintTime)
	}

	return r
}

// Finalize ensures there no nil pointers.
func (c *DogStatsDConfig) Finalize() {
	if c == nil {
		return
	}

	d := DefaultDogStatsDConfig()

	if c.Address == nil {
		c.Address = d.Address
	}

	if c.Period == nil {
		c.Period = d.Period
	}
}

// Copy returns a deep copy of this configuration.
func (c *DogStatsDConfig) Copy() *DogStatsDConfig {
	if c == nil {
		return nil
	}

	return &DogStatsDConfig{
		Address: StringCopy(c.Address),
		Period:  TimeDurationCopy(c.Period),
	}
}

// Merge combines all values in this configuration with the values in the other
// configuration, with values in the other configuration taking precedence.
// Maps and slices are merged, most other values are overwritten. Complex
// structs define their own merge functionality.
func (c *DogStatsDConfig) Merge(o *DogStatsDConfig) *DogStatsDConfig {
	if c == nil {
		if o == nil {
			return nil
		}
		return o.Copy()
	}

	if o == nil {
		return c.Copy()
	}

	r := c.Copy()

	if o.Address != nil {
		r.Address = StringCopy(o.Address)
	}

	if o.Period != nil {
		r.Period = TimeDurationCopy(o.Period)
	}

	return r
}

// DefaultPrometheusConfig returns a configuration that is populated with the
// default values.
func DefaultPrometheusConfig() *PrometheusConfig {
	return &PrometheusConfig{
		Port:        Uint(DefaultPrometheusPort),
		CachePeriod: TimeDuration(time.Duration(10) * time.Second),
	}
}

// Copy returns a deep copy of this configuration.
func (c *PrometheusConfig) Copy() *PrometheusConfig {
	if c == nil {
		return nil
	}

	return &PrometheusConfig{
		Port:        UintCopy(c.Port),
		CachePeriod: TimeDurationCopy(c.CachePeriod),
	}
}

// Merge combines all values in this configuration with the values in the other
// configuration, with values in the other configuration taking precedence.
// Maps and slices are merged, most other values are overwritten. Complex
// structs define their own merge functionality.
func (c *PrometheusConfig) Merge(o *PrometheusConfig) *PrometheusConfig {
	if c == nil {
		if o == nil {
			return nil
		}
		return o.Copy()
	}

	if o == nil {
		return c.Copy()
	}

	r := c.Copy()

	if o.Port != nil {
		r.Port = UintCopy(o.Port)
	}

	if o.CachePeriod != nil {
		r.CachePeriod = TimeDurationCopy(o.CachePeriod)
	}

	return r
}

// Finalize ensures there no nil pointers.
func (c *PrometheusConfig) Finalize() {
	if c == nil {
		return
	}

	d := DefaultPrometheusConfig()

	if c.Port == nil {
		c.Port = d.Port
	}

	if c.CachePeriod == nil {
		c.CachePeriod = d.CachePeriod
	}
}
