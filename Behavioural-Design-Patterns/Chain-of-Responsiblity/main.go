package main

import "fmt"

//////////////// Interface ///////////////////////

type department interface {
	execute(*patient)
	setNext(department)
}

//////////////// Departments /////////////////////

///////////////// Reception //////////////////////

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.receptionDone {
		fmt.Println("Patient registration already done.")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient.")
	p.receptionDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

///////////////// Doctor /////////////////////////

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorDone {
		fmt.Println("Doctor checkup already done.")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient.")
	p.doctorDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

///////////////// Medical ////////////////////////

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicalDone {
		fmt.Println("Medicine already given to patient.")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient.")
	p.medicalDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

////////////////// Cashier ///////////////////////

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.cashierDone {
		fmt.Println("Payment Done.")
	}
	fmt.Println("Cashier getting money from patient patient.")
}

func (c *cashier) setNext(next department) {
	c.next = next
}

// We send a patient through the departments.
type patient struct {
	name          string
	receptionDone bool
	doctorDone    bool
	medicalDone   bool
	cashierDone   bool
}

func main() {
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "Sick Bastard"}
	reception.execute(patient)
}
