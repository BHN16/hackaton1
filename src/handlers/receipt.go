package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
)

func GetReceipts(w http.ResponseWriter, r *http.Request) {

	var receipts []models.Receipt

	role, err := processCookie(r)

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"response": "No cookie"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"response": "Bad request"})
		return
	}

	switch role {
	case 1:
		bd.DB.Find(&receipts)

		json.NewEncoder(w).Encode(&receipts)
	case 2:
		c, err := r.Cookie("token")
		if err != nil {
			return
		}
		bd.DB.Find(&receipts, "Employee_Refer = ?", getEmail(c.Value))
		json.NewEncoder(w).Encode(&receipts)
		return
	case 3:
		c, err := r.Cookie("token")
		if err != nil {
			return
		}
		bd.DB.Find(&receipts, "Patient_Refer = ?", getEmail(c.Value))
		json.NewEncoder(w).Encode(&receipts)
		return
	default:
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})
		return

	}

}

func PostReceipt(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		return
	}
        var err_demo bool = true

	var receipt models.Receipt
	var employee models.Employee
	var patient models.Patient
	var medicine models.Medicine



        if err_demo {
                var tempreceipt models.TempReceipt
                err3 := json.NewDecoder(r.Body).Decode(&tempreceipt)
                if err3 == nil{
                        w.WriteHeader(http.StatusBadRequest)
                        dataLog, _ := json.Marshal(tempreceipt)
                        //dataLog, _ := r.Body
                        ErrorLogger.Println("Transaction Error", string(dataLog))
                        json.NewEncoder(w).Encode(map[string]string{"response": "Transaction Error"})
                        return
                }
        }





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
