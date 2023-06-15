package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdLeaseUnitVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_lease_unit_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_lease_unit_version

	for rows.Next() {
		var prod_lease_unit_version dto.Prod_lease_unit_version
		if err := rows.Scan(&prod_lease_unit_version.Lease_unit_id, &prod_lease_unit_version.Source, &prod_lease_unit_version.Acreage, &prod_lease_unit_version.Acreage_ouom, &prod_lease_unit_version.Active_ind, &prod_lease_unit_version.Current_operator, &prod_lease_unit_version.Current_status_date, &prod_lease_unit_version.Effective_date, &prod_lease_unit_version.Expiry_date, &prod_lease_unit_version.Field_id, &prod_lease_unit_version.Government_lease_unit_id, &prod_lease_unit_version.Land_right_id, &prod_lease_unit_version.Land_right_subtype, &prod_lease_unit_version.Lease_unit_long_name, &prod_lease_unit_version.Lease_unit_short_name, &prod_lease_unit_version.Lease_unit_status, &prod_lease_unit_version.Lease_unit_type, &prod_lease_unit_version.Pool_id, &prod_lease_unit_version.Ppdm_guid, &prod_lease_unit_version.Remark, &prod_lease_unit_version.Strat_name_set_id, &prod_lease_unit_version.Strat_unit_id, &prod_lease_unit_version.Row_changed_by, &prod_lease_unit_version.Row_changed_date, &prod_lease_unit_version.Row_created_by, &prod_lease_unit_version.Row_created_date, &prod_lease_unit_version.Row_effective_date, &prod_lease_unit_version.Row_expiry_date, &prod_lease_unit_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_lease_unit_version to result
		result = append(result, prod_lease_unit_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdLeaseUnitVersion(c *fiber.Ctx) error {
	var prod_lease_unit_version dto.Prod_lease_unit_version

	setDefaults(&prod_lease_unit_version)

	if err := json.Unmarshal(c.Body(), &prod_lease_unit_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_lease_unit_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	prod_lease_unit_version.Row_created_date = formatDate(prod_lease_unit_version.Row_created_date)
	prod_lease_unit_version.Row_changed_date = formatDate(prod_lease_unit_version.Row_changed_date)
	prod_lease_unit_version.Current_status_date = formatDateString(prod_lease_unit_version.Current_status_date)
	prod_lease_unit_version.Effective_date = formatDateString(prod_lease_unit_version.Effective_date)
	prod_lease_unit_version.Expiry_date = formatDateString(prod_lease_unit_version.Expiry_date)
	prod_lease_unit_version.Row_effective_date = formatDateString(prod_lease_unit_version.Row_effective_date)
	prod_lease_unit_version.Row_expiry_date = formatDateString(prod_lease_unit_version.Row_expiry_date)






	rows, err := stmt.Exec(prod_lease_unit_version.Lease_unit_id, prod_lease_unit_version.Source, prod_lease_unit_version.Acreage, prod_lease_unit_version.Acreage_ouom, prod_lease_unit_version.Active_ind, prod_lease_unit_version.Current_operator, prod_lease_unit_version.Current_status_date, prod_lease_unit_version.Effective_date, prod_lease_unit_version.Expiry_date, prod_lease_unit_version.Field_id, prod_lease_unit_version.Government_lease_unit_id, prod_lease_unit_version.Land_right_id, prod_lease_unit_version.Land_right_subtype, prod_lease_unit_version.Lease_unit_long_name, prod_lease_unit_version.Lease_unit_short_name, prod_lease_unit_version.Lease_unit_status, prod_lease_unit_version.Lease_unit_type, prod_lease_unit_version.Pool_id, prod_lease_unit_version.Ppdm_guid, prod_lease_unit_version.Remark, prod_lease_unit_version.Strat_name_set_id, prod_lease_unit_version.Strat_unit_id, prod_lease_unit_version.Row_changed_by, prod_lease_unit_version.Row_changed_date, prod_lease_unit_version.Row_created_by, prod_lease_unit_version.Row_created_date, prod_lease_unit_version.Row_effective_date, prod_lease_unit_version.Row_expiry_date, prod_lease_unit_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdLeaseUnitVersion(c *fiber.Ctx) error {
	var prod_lease_unit_version dto.Prod_lease_unit_version

	setDefaults(&prod_lease_unit_version)

	if err := json.Unmarshal(c.Body(), &prod_lease_unit_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_lease_unit_version.Lease_unit_id = id

    if prod_lease_unit_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_lease_unit_version set source = :1, acreage = :2, acreage_ouom = :3, active_ind = :4, current_operator = :5, current_status_date = :6, effective_date = :7, expiry_date = :8, field_id = :9, government_lease_unit_id = :10, land_right_id = :11, land_right_subtype = :12, lease_unit_long_name = :13, lease_unit_short_name = :14, lease_unit_status = :15, lease_unit_type = :16, pool_id = :17, ppdm_guid = :18, remark = :19, strat_name_set_id = :20, strat_unit_id = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where lease_unit_id = :29")
	if err != nil {
		return err
	}

	prod_lease_unit_version.Row_changed_date = formatDate(prod_lease_unit_version.Row_changed_date)
	prod_lease_unit_version.Current_status_date = formatDateString(prod_lease_unit_version.Current_status_date)
	prod_lease_unit_version.Effective_date = formatDateString(prod_lease_unit_version.Effective_date)
	prod_lease_unit_version.Expiry_date = formatDateString(prod_lease_unit_version.Expiry_date)
	prod_lease_unit_version.Row_effective_date = formatDateString(prod_lease_unit_version.Row_effective_date)
	prod_lease_unit_version.Row_expiry_date = formatDateString(prod_lease_unit_version.Row_expiry_date)






	rows, err := stmt.Exec(prod_lease_unit_version.Source, prod_lease_unit_version.Acreage, prod_lease_unit_version.Acreage_ouom, prod_lease_unit_version.Active_ind, prod_lease_unit_version.Current_operator, prod_lease_unit_version.Current_status_date, prod_lease_unit_version.Effective_date, prod_lease_unit_version.Expiry_date, prod_lease_unit_version.Field_id, prod_lease_unit_version.Government_lease_unit_id, prod_lease_unit_version.Land_right_id, prod_lease_unit_version.Land_right_subtype, prod_lease_unit_version.Lease_unit_long_name, prod_lease_unit_version.Lease_unit_short_name, prod_lease_unit_version.Lease_unit_status, prod_lease_unit_version.Lease_unit_type, prod_lease_unit_version.Pool_id, prod_lease_unit_version.Ppdm_guid, prod_lease_unit_version.Remark, prod_lease_unit_version.Strat_name_set_id, prod_lease_unit_version.Strat_unit_id, prod_lease_unit_version.Row_changed_by, prod_lease_unit_version.Row_changed_date, prod_lease_unit_version.Row_created_by, prod_lease_unit_version.Row_effective_date, prod_lease_unit_version.Row_expiry_date, prod_lease_unit_version.Row_quality, prod_lease_unit_version.Lease_unit_id)
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

func PatchProdLeaseUnitVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_lease_unit_version set "
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
		} else if key == "current_status_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteProdLeaseUnitVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_lease_unit_version dto.Prod_lease_unit_version
	prod_lease_unit_version.Lease_unit_id = id

	stmt, err := db.Prepare("delete from prod_lease_unit_version where lease_unit_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_lease_unit_version.Lease_unit_id)
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


