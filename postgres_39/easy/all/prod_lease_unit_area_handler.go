package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdLeaseUnitArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_lease_unit_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_lease_unit_area

	for rows.Next() {
		var prod_lease_unit_area dto.Prod_lease_unit_area
		if err := rows.Scan(&prod_lease_unit_area.Lease_unit_id, &prod_lease_unit_area.Source, &prod_lease_unit_area.Area_id, &prod_lease_unit_area.Area_type, &prod_lease_unit_area.Active_ind, &prod_lease_unit_area.Effective_date, &prod_lease_unit_area.Expiry_date, &prod_lease_unit_area.Ppdm_guid, &prod_lease_unit_area.Remark, &prod_lease_unit_area.Row_changed_by, &prod_lease_unit_area.Row_changed_date, &prod_lease_unit_area.Row_created_by, &prod_lease_unit_area.Row_created_date, &prod_lease_unit_area.Row_effective_date, &prod_lease_unit_area.Row_expiry_date, &prod_lease_unit_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_lease_unit_area to result
		result = append(result, prod_lease_unit_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdLeaseUnitArea(c *fiber.Ctx) error {
	var prod_lease_unit_area dto.Prod_lease_unit_area

	setDefaults(&prod_lease_unit_area)

	if err := json.Unmarshal(c.Body(), &prod_lease_unit_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_lease_unit_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	prod_lease_unit_area.Row_created_date = formatDate(prod_lease_unit_area.Row_created_date)
	prod_lease_unit_area.Row_changed_date = formatDate(prod_lease_unit_area.Row_changed_date)
	prod_lease_unit_area.Effective_date = formatDateString(prod_lease_unit_area.Effective_date)
	prod_lease_unit_area.Expiry_date = formatDateString(prod_lease_unit_area.Expiry_date)
	prod_lease_unit_area.Row_effective_date = formatDateString(prod_lease_unit_area.Row_effective_date)
	prod_lease_unit_area.Row_expiry_date = formatDateString(prod_lease_unit_area.Row_expiry_date)






	rows, err := stmt.Exec(prod_lease_unit_area.Lease_unit_id, prod_lease_unit_area.Source, prod_lease_unit_area.Area_id, prod_lease_unit_area.Area_type, prod_lease_unit_area.Active_ind, prod_lease_unit_area.Effective_date, prod_lease_unit_area.Expiry_date, prod_lease_unit_area.Ppdm_guid, prod_lease_unit_area.Remark, prod_lease_unit_area.Row_changed_by, prod_lease_unit_area.Row_changed_date, prod_lease_unit_area.Row_created_by, prod_lease_unit_area.Row_created_date, prod_lease_unit_area.Row_effective_date, prod_lease_unit_area.Row_expiry_date, prod_lease_unit_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdLeaseUnitArea(c *fiber.Ctx) error {
	var prod_lease_unit_area dto.Prod_lease_unit_area

	setDefaults(&prod_lease_unit_area)

	if err := json.Unmarshal(c.Body(), &prod_lease_unit_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_lease_unit_area.Lease_unit_id = id

    if prod_lease_unit_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_lease_unit_area set source = :1, area_id = :2, area_type = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where lease_unit_id = :16")
	if err != nil {
		return err
	}

	prod_lease_unit_area.Row_changed_date = formatDate(prod_lease_unit_area.Row_changed_date)
	prod_lease_unit_area.Effective_date = formatDateString(prod_lease_unit_area.Effective_date)
	prod_lease_unit_area.Expiry_date = formatDateString(prod_lease_unit_area.Expiry_date)
	prod_lease_unit_area.Row_effective_date = formatDateString(prod_lease_unit_area.Row_effective_date)
	prod_lease_unit_area.Row_expiry_date = formatDateString(prod_lease_unit_area.Row_expiry_date)






	rows, err := stmt.Exec(prod_lease_unit_area.Source, prod_lease_unit_area.Area_id, prod_lease_unit_area.Area_type, prod_lease_unit_area.Active_ind, prod_lease_unit_area.Effective_date, prod_lease_unit_area.Expiry_date, prod_lease_unit_area.Ppdm_guid, prod_lease_unit_area.Remark, prod_lease_unit_area.Row_changed_by, prod_lease_unit_area.Row_changed_date, prod_lease_unit_area.Row_created_by, prod_lease_unit_area.Row_effective_date, prod_lease_unit_area.Row_expiry_date, prod_lease_unit_area.Row_quality, prod_lease_unit_area.Lease_unit_id)
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

func PatchProdLeaseUnitArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_lease_unit_area set "
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
	query += " where lease_unit_id = :id"

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

func DeleteProdLeaseUnitArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_lease_unit_area dto.Prod_lease_unit_area
	prod_lease_unit_area.Lease_unit_id = id

	stmt, err := db.Prepare("delete from prod_lease_unit_area where lease_unit_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_lease_unit_area.Lease_unit_id)
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


