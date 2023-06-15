package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdString(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_string")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_string

	for rows.Next() {
		var prod_string dto.Prod_string
		if err := rows.Scan(&prod_string.Uwi, &prod_string.Source, &prod_string.String_id, &prod_string.Active_ind, &prod_string.Base_depth, &prod_string.Base_depth_ouom, &prod_string.Business_associate_id, &prod_string.Commingled_ind, &prod_string.Current_status, &prod_string.Current_status_date, &prod_string.Effective_date, &prod_string.Expiry_date, &prod_string.Facility_id, &prod_string.Facility_type, &prod_string.Field_id, &prod_string.Government_string_id, &prod_string.Lease_unit_id, &prod_string.On_injection_date, &prod_string.On_production_date, &prod_string.Plot_symbol, &prod_string.Pool_id, &prod_string.Ppdm_guid, &prod_string.Prod_string_tvd, &prod_string.Prod_string_tvd_ouom, &prod_string.Prod_string_type, &prod_string.Profile_type, &prod_string.Remark, &prod_string.Status_type, &prod_string.Strat_name_set_id, &prod_string.Strat_unit_id, &prod_string.Tax_credit_code, &prod_string.Top_depth, &prod_string.Top_depth_ouom, &prod_string.Total_depth, &prod_string.Total_depth_ouom, &prod_string.Row_changed_by, &prod_string.Row_changed_date, &prod_string.Row_created_by, &prod_string.Row_created_date, &prod_string.Row_effective_date, &prod_string.Row_expiry_date, &prod_string.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_string to result
		result = append(result, prod_string)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdString(c *fiber.Ctx) error {
	var prod_string dto.Prod_string

	setDefaults(&prod_string)

	if err := json.Unmarshal(c.Body(), &prod_string); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_string values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	prod_string.Row_created_date = formatDate(prod_string.Row_created_date)
	prod_string.Row_changed_date = formatDate(prod_string.Row_changed_date)
	prod_string.Current_status_date = formatDateString(prod_string.Current_status_date)
	prod_string.Effective_date = formatDateString(prod_string.Effective_date)
	prod_string.Expiry_date = formatDateString(prod_string.Expiry_date)
	prod_string.On_injection_date = formatDateString(prod_string.On_injection_date)
	prod_string.On_production_date = formatDateString(prod_string.On_production_date)
	prod_string.Row_effective_date = formatDateString(prod_string.Row_effective_date)
	prod_string.Row_expiry_date = formatDateString(prod_string.Row_expiry_date)






	rows, err := stmt.Exec(prod_string.Uwi, prod_string.Source, prod_string.String_id, prod_string.Active_ind, prod_string.Base_depth, prod_string.Base_depth_ouom, prod_string.Business_associate_id, prod_string.Commingled_ind, prod_string.Current_status, prod_string.Current_status_date, prod_string.Effective_date, prod_string.Expiry_date, prod_string.Facility_id, prod_string.Facility_type, prod_string.Field_id, prod_string.Government_string_id, prod_string.Lease_unit_id, prod_string.On_injection_date, prod_string.On_production_date, prod_string.Plot_symbol, prod_string.Pool_id, prod_string.Ppdm_guid, prod_string.Prod_string_tvd, prod_string.Prod_string_tvd_ouom, prod_string.Prod_string_type, prod_string.Profile_type, prod_string.Remark, prod_string.Status_type, prod_string.Strat_name_set_id, prod_string.Strat_unit_id, prod_string.Tax_credit_code, prod_string.Top_depth, prod_string.Top_depth_ouom, prod_string.Total_depth, prod_string.Total_depth_ouom, prod_string.Row_changed_by, prod_string.Row_changed_date, prod_string.Row_created_by, prod_string.Row_created_date, prod_string.Row_effective_date, prod_string.Row_expiry_date, prod_string.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdString(c *fiber.Ctx) error {
	var prod_string dto.Prod_string

	setDefaults(&prod_string)

	if err := json.Unmarshal(c.Body(), &prod_string); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_string.Uwi = id

    if prod_string.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_string set source = :1, string_id = :2, active_ind = :3, base_depth = :4, base_depth_ouom = :5, business_associate_id = :6, commingled_ind = :7, current_status = :8, current_status_date = :9, effective_date = :10, expiry_date = :11, facility_id = :12, facility_type = :13, field_id = :14, government_string_id = :15, lease_unit_id = :16, on_injection_date = :17, on_production_date = :18, plot_symbol = :19, pool_id = :20, ppdm_guid = :21, prod_string_tvd = :22, prod_string_tvd_ouom = :23, prod_string_type = :24, profile_type = :25, remark = :26, status_type = :27, strat_name_set_id = :28, strat_unit_id = :29, tax_credit_code = :30, top_depth = :31, top_depth_ouom = :32, total_depth = :33, total_depth_ouom = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where uwi = :42")
	if err != nil {
		return err
	}

	prod_string.Row_changed_date = formatDate(prod_string.Row_changed_date)
	prod_string.Current_status_date = formatDateString(prod_string.Current_status_date)
	prod_string.Effective_date = formatDateString(prod_string.Effective_date)
	prod_string.Expiry_date = formatDateString(prod_string.Expiry_date)
	prod_string.On_injection_date = formatDateString(prod_string.On_injection_date)
	prod_string.On_production_date = formatDateString(prod_string.On_production_date)
	prod_string.Row_effective_date = formatDateString(prod_string.Row_effective_date)
	prod_string.Row_expiry_date = formatDateString(prod_string.Row_expiry_date)






	rows, err := stmt.Exec(prod_string.Source, prod_string.String_id, prod_string.Active_ind, prod_string.Base_depth, prod_string.Base_depth_ouom, prod_string.Business_associate_id, prod_string.Commingled_ind, prod_string.Current_status, prod_string.Current_status_date, prod_string.Effective_date, prod_string.Expiry_date, prod_string.Facility_id, prod_string.Facility_type, prod_string.Field_id, prod_string.Government_string_id, prod_string.Lease_unit_id, prod_string.On_injection_date, prod_string.On_production_date, prod_string.Plot_symbol, prod_string.Pool_id, prod_string.Ppdm_guid, prod_string.Prod_string_tvd, prod_string.Prod_string_tvd_ouom, prod_string.Prod_string_type, prod_string.Profile_type, prod_string.Remark, prod_string.Status_type, prod_string.Strat_name_set_id, prod_string.Strat_unit_id, prod_string.Tax_credit_code, prod_string.Top_depth, prod_string.Top_depth_ouom, prod_string.Total_depth, prod_string.Total_depth_ouom, prod_string.Row_changed_by, prod_string.Row_changed_date, prod_string.Row_created_by, prod_string.Row_effective_date, prod_string.Row_expiry_date, prod_string.Row_quality, prod_string.Uwi)
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

func PatchProdString(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_string set "
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
		} else if key == "current_status_date" || key == "effective_date" || key == "expiry_date" || key == "on_injection_date" || key == "on_production_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteProdString(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_string dto.Prod_string
	prod_string.Uwi = id

	stmt, err := db.Prepare("delete from prod_string where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_string.Uwi)
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


