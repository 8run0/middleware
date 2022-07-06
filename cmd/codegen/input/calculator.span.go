package calculator

var _ calculator = &calculatorSpanner{}

type calculator interface {
	add(x int64, y int64) (z int64)
	sub(x int64, y int64) (z int64)
	multiply(x int64, y int64) (z int64)
	divide(x int64, y int64) (z int64)
}

type calculatorSpanner struct {
	*OTELTools
	next calculator
}

func (s *calculatorSpanner) add(x int64, y int64) (z int64) {
	ctx, span := s.Tracer.Start(s.Ctx, "calculator_add")
	s.Ctx = ctx
	defer span.End()
	return s.next.add(x, y)
}

func (s *calculatorSpanner) sub(x int64, y int64) (z int64) {
	ctx, span := s.Tracer.Start(s.Ctx, "calculator_sub")
	s.Ctx = ctx
	defer span.End()
	return s.next.sub(x, y)
}

func (s *calculatorSpanner) multiply(x int64, y int64) (z int64) {
	ctx, span := s.Tracer.Start(s.Ctx, "calculator_multiply")
	s.Ctx = ctx
	defer span.End()
	return s.next.multiply(x, y)
}

func (s *calculatorSpanner) divide(x int64, y int64) (z int64) {
	ctx, span := s.Tracer.Start(s.Ctx, "calculator_divide")
	s.Ctx = ctx
	defer span.End()
	return s.next.divide(x, y)
}
