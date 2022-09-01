package visitor

type Visitor interface {
	visitPSU(PSU)
	visitCPU(CPU)
}

type PricingVisitor struct {
	totalCost int
}

func (p *PricingVisitor) visitPSU(psu PSU) {
	p.totalCost += psu.getPrice()
}

func (p *PricingVisitor) visitCPU(cpu CPU) {
	p.totalCost += cpu.getPrice()
}

type PowerRequirementVisitor struct {
	totalPowerConsumption int
}

func (p *PowerRequirementVisitor) visitPSU(psu PSU) {
	p.totalPowerConsumption += psu.getPower()
}

func (p *PowerRequirementVisitor) visitCPU(cpu CPU) {
	p.totalPowerConsumption += cpu.getPower()
}
