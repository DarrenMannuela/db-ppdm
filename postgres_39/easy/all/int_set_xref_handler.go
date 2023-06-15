package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetIntSetXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from int_set_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Int_set_xref

	for rows.Next() {
		var int_set_xref dto.Int_set_xref
		if err := rows.Scan(&int_set_xref.Interest_set_id, &int_set_xref.Interest_set_seq_no, &int_set_xref.Interest_set_id_2, &int_set_xref.Interest_set_seq_no_2, &int_set_xref.Int_set_xref_obs_no, &int_set_xref.Active_ind, &int_set_xref.Contract_id, &int_set_xref.Effective_date, &int_set_xref.Expiry_date, &int_set_xref.Int_set_xref_type, &int_set_xref.Partner_ba_id, &int_set_xref.Partner_ba_id_2, &int_set_xref.Partner_obs_no, &int_set_xref.Partner_obs_no_2, &int_set_xref.Ppdm_guid, &int_set_xref.Provision_id, &int_set_xref.Remark, &int_set_xref.Source, &int_set_xref.Row_changed_by, &int_set_xref.Row_changed_date, &int_set_xref.Row_created_by, &int_set_xref.Row_created_date, &int_set_xref.Row_effective_date, &int_set_xref.Row_expiry_date, &int_set_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append int_set_xref to result
		result = append(result, int_set_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetIntSetXref(c *fiber.Ctx) error {
	var int_set_xref dto.Int_set_xref

	setDefaults(&int_set_xref)

	if err := json.Unmarshal(c.Body(), &int_set_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into int_set_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	int_set_xref.Row_created_date = formatDate(int_set_xref.Row_created_date)
	int_set_xref.Row_changed_date = formatDate(int_set_xref.Row_changed_date)
	int_set_xref.Effective_date = formatDateString(int_set_xref.Effective_date)
	int_set_xref.Expiry_date = formatDateString(int_set_xref.Expiry_date)
	int_set_xref.Row_effective_date = formatDateString(int_set_xref.Row_effective_date)
	int_set_xref.Row_expiry_date = formatDateString(int_set_xref.Row_expiry_date)






	rows, err := stmt.Exec(int_set_xref.Interest_set_id, int_set_xref.Interest_set_seq_no, int_set_xref.Interest_set_id_2, int_set_xref.Interest_set_seq_no_2, int_set_xref.Int_set_xref_obs_no, int_set_xref.Active_ind, int_set_xref.Contract_id, int_set_xref.Effective_date, int_set_xref.Expiry_date, int_set_xref.Int_set_xref_type, int_set_xref.Partner_ba_id, int_set_xref.Partner_ba_id_2, int_set_xref.Partner_obs_no, int_set_xref.Partner_obs_no_2, int_set_xref.Ppdm_guid, int_set_xref.Provision_id, int_set_xref.Remark, int_set_xref.Source, int_set_xref.Row_changed_by, int_set_xref.Row_changed_date, int_set_xref.Row_created_by, int_set_xref.Row_created_date, int_set_xref.Row_effective_date, int_set_xref.Row_expiry_date, int_set_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateIntSetXref(c *fiber.Ctx) error {
	var int_set_xref dto.Int_set_xref

	setDefaults(&int_set_xref)

	if err := json.Unmarshal(c.Body(), &int_set_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	int_set_xref.Interest_set_id = id

    if int_set_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update int_set_xref set interest_set_seq_no = :1, interest_set_id_2 = :2, interest_set_seq_no_2 = :3, int_set_xref_obs_no = :4, active_ind = :5, contract_id = :6, effective_date = :7, expiry_date = :8, int_set_xref_type = :9, partner_ba_id = :10, partner_ba_id_2 = :11, partner_obs_no = :12, partner_obs_no_2 = :13, ppdm_guid = :14, provision_id = :15, remark = :16, source = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where interest_set_id = :25")
	if err != nil {
		return err
	}

	int_set_xref.Row_changed_date = formatDate(int_set_xref.Row_changed_date)
	int_set_xref.Effective_date = formatDateString(int_set_xref.Effective_date)
	int_set_xref.Expiry_date = formatDateString(int_set_xref.Expiry_date)
	int_set_xref.Row_effective_date = formatDateString(int_set_xref.Row_effective_date)
	int_set_xref.Row_expiry_date = formatDateString(int_set_xref.Row_expiry_date)






	rows, err := stmt.Exec(int_set_xref.Interest_set_seq_no, int_set_xref.Interest_set_id_2, int_set_xref.Interest_set_seq_no_2, int_set_xref.Int_set_xref_obs_no, int_set_xref.Active_ind, int_set_xref.Contract_id, int_set_xref.Effective_date, int_set_xref.Expiry_date, int_set_xref.Int_set_xref_type, int_set_xref.Partner_ba_id, int_set_xref.Partner_ba_id_2, int_set_xref.Partner_obs_no, int_set_xref.Partner_obs_no_2, int_set_xref.Ppdm_guid, int_set_xref.Provision_id, int_set_xref.Remark, int_set_xref.Source, int_set_xref.Row_changed_by, int_set_xref.Row_changed_date, int_set_xref.Row_created_by, int_set_xref.Row_effective_date, int_set_xref.Row_expiry_date, int_set_xref.Row_quality, int_set_xref.Interest_set_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchIntSetXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update int_set_xref set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where interest_set_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteIntSetXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var int_set_xref dto.Int_set_xref
	int_set_xref.Interest_set_id = id

	stmt, err := db.Prepare("delete from int_set_xref where interest_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(int_set_xref.Interest_set_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


