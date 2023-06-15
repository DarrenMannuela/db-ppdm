package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenLeaseUnit(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_lease_unit")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_lease_unit

	for rows.Next() {
		var pden_lease_unit dto.Pden_lease_unit
		if err := rows.Scan(&pden_lease_unit.Pden_subtype, &pden_lease_unit.Pden_id, &pden_lease_unit.Pden_source, &pden_lease_unit.Active_ind, &pden_lease_unit.Effective_date, &pden_lease_unit.Expiry_date, &pden_lease_unit.Facility_id, &pden_lease_unit.Facility_type, &pden_lease_unit.Lease_unit_id, &pden_lease_unit.No_of_gas_wells, &pden_lease_unit.No_of_injection_wells, &pden_lease_unit.No_of_oil_wells, &pden_lease_unit.Ppdm_guid, &pden_lease_unit.Remark, &pden_lease_unit.Row_changed_by, &pden_lease_unit.Row_changed_date, &pden_lease_unit.Row_created_by, &pden_lease_unit.Row_created_date, &pden_lease_unit.Row_effective_date, &pden_lease_unit.Row_expiry_date, &pden_lease_unit.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_lease_unit to result
		result = append(result, pden_lease_unit)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenLeaseUnit(c *fiber.Ctx) error {
	var pden_lease_unit dto.Pden_lease_unit

	setDefaults(&pden_lease_unit)

	if err := json.Unmarshal(c.Body(), &pden_lease_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_lease_unit values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	pden_lease_unit.Row_created_date = formatDate(pden_lease_unit.Row_created_date)
	pden_lease_unit.Row_changed_date = formatDate(pden_lease_unit.Row_changed_date)
	pden_lease_unit.Effective_date = formatDateString(pden_lease_unit.Effective_date)
	pden_lease_unit.Expiry_date = formatDateString(pden_lease_unit.Expiry_date)
	pden_lease_unit.Row_effective_date = formatDateString(pden_lease_unit.Row_effective_date)
	pden_lease_unit.Row_expiry_date = formatDateString(pden_lease_unit.Row_expiry_date)






	rows, err := stmt.Exec(pden_lease_unit.Pden_subtype, pden_lease_unit.Pden_id, pden_lease_unit.Pden_source, pden_lease_unit.Active_ind, pden_lease_unit.Effective_date, pden_lease_unit.Expiry_date, pden_lease_unit.Facility_id, pden_lease_unit.Facility_type, pden_lease_unit.Lease_unit_id, pden_lease_unit.No_of_gas_wells, pden_lease_unit.No_of_injection_wells, pden_lease_unit.No_of_oil_wells, pden_lease_unit.Ppdm_guid, pden_lease_unit.Remark, pden_lease_unit.Row_changed_by, pden_lease_unit.Row_changed_date, pden_lease_unit.Row_created_by, pden_lease_unit.Row_created_date, pden_lease_unit.Row_effective_date, pden_lease_unit.Row_expiry_date, pden_lease_unit.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenLeaseUnit(c *fiber.Ctx) error {
	var pden_lease_unit dto.Pden_lease_unit

	setDefaults(&pden_lease_unit)

	if err := json.Unmarshal(c.Body(), &pden_lease_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_lease_unit.Pden_subtype = id

    if pden_lease_unit.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_lease_unit set pden_id = :1, pden_source = :2, active_ind = :3, effective_date = :4, expiry_date = :5, facility_id = :6, facility_type = :7, lease_unit_id = :8, no_of_gas_wells = :9, no_of_injection_wells = :10, no_of_oil_wells = :11, ppdm_guid = :12, remark = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where pden_subtype = :21")
	if err != nil {
		return err
	}

	pden_lease_unit.Row_changed_date = formatDate(pden_lease_unit.Row_changed_date)
	pden_lease_unit.Effective_date = formatDateString(pden_lease_unit.Effective_date)
	pden_lease_unit.Expiry_date = formatDateString(pden_lease_unit.Expiry_date)
	pden_lease_unit.Row_effective_date = formatDateString(pden_lease_unit.Row_effective_date)
	pden_lease_unit.Row_expiry_date = formatDateString(pden_lease_unit.Row_expiry_date)






	rows, err := stmt.Exec(pden_lease_unit.Pden_id, pden_lease_unit.Pden_source, pden_lease_unit.Active_ind, pden_lease_unit.Effective_date, pden_lease_unit.Expiry_date, pden_lease_unit.Facility_id, pden_lease_unit.Facility_type, pden_lease_unit.Lease_unit_id, pden_lease_unit.No_of_gas_wells, pden_lease_unit.No_of_injection_wells, pden_lease_unit.No_of_oil_wells, pden_lease_unit.Ppdm_guid, pden_lease_unit.Remark, pden_lease_unit.Row_changed_by, pden_lease_unit.Row_changed_date, pden_lease_unit.Row_created_by, pden_lease_unit.Row_effective_date, pden_lease_unit.Row_expiry_date, pden_lease_unit.Row_quality, pden_lease_unit.Pden_subtype)
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

func PatchPdenLeaseUnit(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_lease_unit set "
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

func DeletePdenLeaseUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_lease_unit dto.Pden_lease_unit
	pden_lease_unit.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_lease_unit where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_lease_unit.Pden_subtype)
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


