package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenLandRight(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_land_right")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_land_right

	for rows.Next() {
		var pden_land_right dto.Pden_land_right
		if err := rows.Scan(&pden_land_right.Pden_subtype, &pden_land_right.Pden_id, &pden_land_right.Pden_source, &pden_land_right.Active_ind, &pden_land_right.Effective_date, &pden_land_right.Expiry_date, &pden_land_right.Facility_id, &pden_land_right.Facility_type, &pden_land_right.Land_right_id, &pden_land_right.Land_right_subtype, &pden_land_right.No_of_gas_wells, &pden_land_right.No_of_injection_wells, &pden_land_right.No_of_oil_wells, &pden_land_right.Ppdm_guid, &pden_land_right.Remark, &pden_land_right.Row_changed_by, &pden_land_right.Row_changed_date, &pden_land_right.Row_created_by, &pden_land_right.Row_created_date, &pden_land_right.Row_effective_date, &pden_land_right.Row_expiry_date, &pden_land_right.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_land_right to result
		result = append(result, pden_land_right)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenLandRight(c *fiber.Ctx) error {
	var pden_land_right dto.Pden_land_right

	setDefaults(&pden_land_right)

	if err := json.Unmarshal(c.Body(), &pden_land_right); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_land_right values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	pden_land_right.Row_created_date = formatDate(pden_land_right.Row_created_date)
	pden_land_right.Row_changed_date = formatDate(pden_land_right.Row_changed_date)
	pden_land_right.Effective_date = formatDateString(pden_land_right.Effective_date)
	pden_land_right.Expiry_date = formatDateString(pden_land_right.Expiry_date)
	pden_land_right.Row_effective_date = formatDateString(pden_land_right.Row_effective_date)
	pden_land_right.Row_expiry_date = formatDateString(pden_land_right.Row_expiry_date)






	rows, err := stmt.Exec(pden_land_right.Pden_subtype, pden_land_right.Pden_id, pden_land_right.Pden_source, pden_land_right.Active_ind, pden_land_right.Effective_date, pden_land_right.Expiry_date, pden_land_right.Facility_id, pden_land_right.Facility_type, pden_land_right.Land_right_id, pden_land_right.Land_right_subtype, pden_land_right.No_of_gas_wells, pden_land_right.No_of_injection_wells, pden_land_right.No_of_oil_wells, pden_land_right.Ppdm_guid, pden_land_right.Remark, pden_land_right.Row_changed_by, pden_land_right.Row_changed_date, pden_land_right.Row_created_by, pden_land_right.Row_created_date, pden_land_right.Row_effective_date, pden_land_right.Row_expiry_date, pden_land_right.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenLandRight(c *fiber.Ctx) error {
	var pden_land_right dto.Pden_land_right

	setDefaults(&pden_land_right)

	if err := json.Unmarshal(c.Body(), &pden_land_right); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_land_right.Pden_subtype = id

    if pden_land_right.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_land_right set pden_id = :1, pden_source = :2, active_ind = :3, effective_date = :4, expiry_date = :5, facility_id = :6, facility_type = :7, land_right_id = :8, land_right_subtype = :9, no_of_gas_wells = :10, no_of_injection_wells = :11, no_of_oil_wells = :12, ppdm_guid = :13, remark = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where pden_subtype = :22")
	if err != nil {
		return err
	}

	pden_land_right.Row_changed_date = formatDate(pden_land_right.Row_changed_date)
	pden_land_right.Effective_date = formatDateString(pden_land_right.Effective_date)
	pden_land_right.Expiry_date = formatDateString(pden_land_right.Expiry_date)
	pden_land_right.Row_effective_date = formatDateString(pden_land_right.Row_effective_date)
	pden_land_right.Row_expiry_date = formatDateString(pden_land_right.Row_expiry_date)






	rows, err := stmt.Exec(pden_land_right.Pden_id, pden_land_right.Pden_source, pden_land_right.Active_ind, pden_land_right.Effective_date, pden_land_right.Expiry_date, pden_land_right.Facility_id, pden_land_right.Facility_type, pden_land_right.Land_right_id, pden_land_right.Land_right_subtype, pden_land_right.No_of_gas_wells, pden_land_right.No_of_injection_wells, pden_land_right.No_of_oil_wells, pden_land_right.Ppdm_guid, pden_land_right.Remark, pden_land_right.Row_changed_by, pden_land_right.Row_changed_date, pden_land_right.Row_created_by, pden_land_right.Row_effective_date, pden_land_right.Row_expiry_date, pden_land_right.Row_quality, pden_land_right.Pden_subtype)
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

func PatchPdenLandRight(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_land_right set "
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
	query += " where pden_subtype = :id"

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

func DeletePdenLandRight(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_land_right dto.Pden_land_right
	pden_land_right.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_land_right where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_land_right.Pden_subtype)
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


