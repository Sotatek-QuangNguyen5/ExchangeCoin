package dto

type Responsive struct {
	Responsive string `json:"Responsive"`
}

func ResponsiveAddSuccess(obj string) *Responsive {

	return &Responsive{

		Responsive: "Add " + obj + " Success!!!",
	}
}

func ResponsiveCreateSuccess(obj string) *Responsive {

	return &Responsive{

		Responsive: "Create " + obj + " Success!!!",
	}
}

func ResponsiveDeleteSuccess(obj string) *Responsive {

	return &Responsive{

		Responsive: "Delete " + obj + " Success!!!",
	}
}

func ResponsiveUpdateSuccess(obj string) *Responsive {

	return &Responsive{

		Responsive: "Update " + obj + " Success!!!",
	}
}