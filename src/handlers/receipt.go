package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
)

func GetReceipts(w http.ResponseWriter, r *http.Request) {

	var receipts []models.Receipt

	bd.DB.Find(&receipts)

	json.NewEncoder(w).Encode(&receipts)

}

func PostReceipt(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		return
	}

	var receipt models.Receipt
	var employee models.Employee
	var patient models.Patient
	var medicine models.Medicine
	err1 := json.NewDecoder(r.Body).Decode(&receipt)

	if err1 != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	employeeEmail := getEmail(c.Value)

	bd.DB.Find(&patient, "Email = ?", receipt.PatientRefer)
	bd.DB.Find(&employee, "Email = ?", employeeEmail)
	bd.DB.Find(&medicine, "ID = ?", receipt.MedicineRefer)

	if patient.Email == "" || patient.Email != receipt.PatientRefer {
		json.NewEncoder(w).Encode(map[string]string{"response": "Null data"})
		return
	}

	if employee.Email == "" || employee.Email != employeeEmail {
		json.NewEncoder(w).Encode(map[string]string{"response": "Null data"})
		return
	}

	if medicine.Name == "" || medicine.ID != receipt.MedicineRefer {
		json.NewEncoder(w).Encode(map[string]string{"response": "Null data"})
		return
	}

	if medicine.Cant < receipt.Cant {
		json.NewEncoder(w).Encode(map[string]string{"response": "Don't have medicines"})
		return
	}

	medicine.Cant -= receipt.Cant

	receipt.EmployeeRefer = employee.Email

	bd.DB.Save(&medicine)

	bd.DB.AutoMigrate(&models.Receipt{})

	bd.DB.Create(&receipt)

	json.NewEncoder(w).Encode(&receipt)

}
