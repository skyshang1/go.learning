package service

import "log"

type CalculateArgs struct {
	A, B int
}

type Calculate int

type CalculateResult int

func (c *Calculate) Multiply(args CalculateArgs, result *CalculateResult) error {
	return Multiply(args, result)
}

func Multiply(args CalculateArgs, result *CalculateResult) error {
	log.Printf("Multiplying %d with %d\n", args.A, args.B)
	*result = CalculateResult(args.A * args.B)
	return nil
}
