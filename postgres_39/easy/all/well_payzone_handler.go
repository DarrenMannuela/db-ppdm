package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPayzone(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_payzone")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_payzone

	for rows.Next() {
		var well_payzone dto.Well_payzone
		if err := rows.Scan(&well_payzone.Uwi, &well_payzone.Source, &well_payzone.Zone_id, &well_payzone.Zone_source, &well_payzone.Payzone_type, &well_payzone.Fluid_type, &well_payzone.Active_ind, &well_payzone.Base_depth, &well_payzone.Base_depth_ouom, &well_payzone.Base_strat_unit_id, &well_payzone.Effective_date, &well_payzone.Expiry_date, &well_payzone.Gas_oil_contact_depth, &well_payzone.Gas_oil_contact_depth_ouom, &well_payzone.Gas_wtr_contact_depth, &well_payzone.Gas_wtr_contact_depth_ouom, &well_payzone.Gross_pay, &well_payzone.Gross_pay_ouom, &well_payzone.Net_pay, &well_payzone.Net_pay_ouom, &well_payzone.Oil_wtr_contact_depth, &well_payzone.Oil_wtr_contact_depth_ouom, &well_payzone.Ppdm_guid, &well_payzone.Remark, &well_payzone.Strat_name_set_id, &well_payzone.Top_depth, &well_payzone.Top_depth_ouom, &well_payzone.Top_strat_unit_id, &well_payzone.Row_changed_by, &well_payzone.Row_changed_date, &well_payzone.Row_created_by, &well_payzone.Row_created_date, &well_payzone.Row_effective_date, &well_payzone.Row_expiry_date, &well_payzone.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_payzone to result
		result = append(result, well_payzone)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPayzone(c *fiber.Ctx) error {
	var well_payzone dto.Well_payzone

	setDefaults(&well_payzone)

	if err := json.Unmarshal(c.Body(), &well_payzone); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_payzone values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	well_payzone.Row_created_date = formatDate(well_payzone.Row_created_date)
	well_payzone.Row_changed_date = formatDate(well_payzone.Row_changed_date)
	well_payzone.Effective_date = formatDateString(well_payzone.Effective_date)
	well_payzone.Expiry_date = formatDateString(well_payzone.Expiry_date)
	well_payzone.Row_effective_date = formatDateString(well_payzone.Row_effective_date)
	well_payzone.Row_expiry_date = formatDateString(well_payzone.Row_expiry_date)






	rows, err := stmt.Exec(well_payzone.Uwi, well_payzone.Source, well_payzone.Zone_id, well_payzone.Zone_source, well_payzone.Payzone_type, well_payzone.Fluid_type, well_payzone.Active_ind, well_payzone.Base_depth, well_payzone.Base_depth_ouom, well_payzone.Base_strat_unit_id, well_payzone.Effective_date, well_payzone.Expiry_date, well_payzone.Gas_oil_contact_depth, well_payzone.Gas_oil_contact_depth_ouom, well_payzone.Gas_wtr_contact_depth, well_payzone.Gas_wtr_contact_depth_ouom, well_payzone.Gross_pay, well_payzone.Gross_pay_ouom, well_payzone.Net_pay, well_payzone.Net_pay_ouom, well_payzone.Oil_wtr_contact_depth, well_payzone.Oil_wtr_contact_depth_ouom, well_payzone.Ppdm_guid, well_payzone.Remark, well_payzone.Strat_name_set_id, well_payzone.Top_depth, well_payzone.Top_depth_ouom, well_payzone.Top_strat_unit_id, well_payzone.Row_changed_by, well_payzone.Row_changed_date, well_payzone.Row_created_by, well_payzone.Row_created_date, well_payzone.Row_effective_date, well_payzone.Row_expiry_date, well_payzone.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPayzone(c *fiber.Ctx) error {
	var well_payzone dto.Well_payzone

	setDefaults(&well_payzone)

	if err := json.Unmarshal(c.Body(), &well_payzone); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_payzone.Uwi = id

    if well_payzone.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_payzone set source = :1, zone_id = :2, zone_source = :3, payzone_type = :4, fluid_type = :5, active_ind = :6, base_depth = :7, base_depth_ouom = :8, base_strat_unit_id = :9, effective_date = :10, expiry_date = :11, gas_oil_contact_depth = :12, gas_oil_contact_depth_ouom = :13, gas_wtr_contact_depth = :14, gas_wtr_contact_depth_ouom = :15, gross_pay = :16, gross_pay_ouom = :17, net_pay = :18, net_pay_ouom = :19, oil_wtr_contact_depth = :20, oil_wtr_contact_depth_ouom = :21, ppdm_guid = :22, remark = :23, strat_name_set_id = :24, top_depth = :25, top_depth_ouom = :26, top_strat_unit_id = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where uwi = :35")
	if err != nil {
		return err
	}

	well_payzone.Row_changed_date = formatDate(well_payzone.Row_changed_date)
	well_payzone.Effective_date = formatDateString(well_payzone.Effective_date)
	well_payzone.Expiry_date = formatDateString(well_payzone.Expiry_date)
	well_payzone.Row_effective_date = formatDateString(well_payzone.Row_effective_date)
	well_payzone.Row_expiry_date = formatDateString(well_payzone.Row_expiry_date)






	rows, err := stmt.Exec(well_payzone.Source, well_payzone.Zone_id, well_payzone.Zone_source, well_payzone.Payzone_type, well_payzone.Fluid_type, well_payzone.Active_ind, well_payzone.Base_depth, well_payzone.Base_depth_ouom, well_payzone.Base_strat_unit_id, well_payzone.Effective_date, well_payzone.Expiry_date, well_payzone.Gas_oil_contact_depth, well_payzone.Gas_oil_contact_depth_ouom, well_payzone.Gas_wtr_contact_depth, well_payzone.Gas_wtr_contact_depth_ouom, well_payzone.Gross_pay, well_payzone.Gross_pay_ouom, well_payzone.Net_pay, well_payzone.Net_pay_ouom, well_payzone.Oil_wtr_contact_depth, well_payzone.Oil_wtr_contact_depth_ouom, well_payzone.Ppdm_guid, well_payzone.Remark, well_payzone.Strat_name_set_id, well_payzone.Top_depth, well_payzone.Top_depth_ouom, well_payzone.Top_strat_unit_id, well_payzone.Row_changed_by, well_payzone.Row_changed_date, well_payzone.Row_created_by, well_payzone.Row_effective_date, well_payzone.Row_expiry_date, well_payzone.Row_quality, well_payzone.Uwi)
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

func PatchWellPayzone(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_payzone set "
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

func DeleteWellPayzone(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_payzone dto.Well_payzone
	well_payzone.Uwi = id

	stmt, err := db.Prepare("delete from well_payzone where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_payzone.Uwi)
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


