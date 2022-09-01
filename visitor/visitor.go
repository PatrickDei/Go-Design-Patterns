package visitor

import "fmt"

func ShowcaseVisitor() {
	fmt.Println("\nVisitor pattern")

	intel := CPU{}
	corsair := PSU{}

	pv := PricingVisitor{totalCost: 0}
	pv.visitPSU(corsair)
	pv.visitCPU(intel)

	fmt.Printf("Total price of components: %v\n", pv.totalCost)

	prv := PowerRequirementVisitor{totalPowerConsumption: 0}
	prv.visitPSU(corsair)
	prv.visitCPU(intel)

	fmt.Printf("Total power consumption of components: %v\n", prv.totalPowerConsumption)
}
