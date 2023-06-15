package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellAirDrill(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_air_drill")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_air_drill

	for rows.Next() {
		var well_air_drill dto.Well_air_drill
		if err := rows.Scan(&well_air_drill.Uwi, &well_air_drill.Source, &well_air_drill.Air_drill_obs_no, &well_air_drill.Active_ind, &well_air_drill.Compressor_count, &well_air_drill.Contractor, &well_air_drill.Effective_date, &well_air_drill.Expiry_date, &well_air_drill.Max_compressor_cap_vol, &well_air_drill.Max_compressor_cap_vol_ouom, &well_air_drill.Ppdm_guid, &well_air_drill.Remark, &well_air_drill.Row_changed_by, &well_air_drill.Row_changed_date, &well_air_drill.Row_created_by, &well_air_drill.Row_created_date, &well_air_drill.Row_effective_date, &well_air_drill.Row_expiry_date, &well_air_drill.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_air_drill to result
		result = append(result, well_air_drill)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellAirDrill(c *fiber.Ctx) error {
	var well_air_drill dto.Well_air_drill

	setDefaults(&well_air_drill)

	if err := json.Unmarshal(c.Body(), &well_air_drill); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_air_drill values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	well_air_drill.Row_created_date = formatDate(well_air_drill.Row_created_date)
	well_air_drill.Row_changed_date = formatDate(well_air_drill.Row_changed_date)
	well_air_drill.Effective_date = formatDateString(well_air_drill.Effective_date)
	well_air_drill.Expiry_date = formatDateString(well_air_drill.Expiry_date)
	well_air_drill.Row_effective_date = formatDateString(well_air_drill.Row_effective_date)
	well_air_drill.Row_expiry_date = formatDateString(well_air_drill.Row_expiry_date)






	rows, err := stmt.Exec(well_air_drill.Uwi, well_air_drill.Source, well_air_drill.Air_drill_obs_no, well_air_drill.Active_ind, well_air_drill.Compressor_count, well_air_drill.Contractor, well_air_drill.Effective_date, well_air_drill.Expiry_date, well_air_drill.Max_compressor_cap_vol, well_air_drill.Max_compressor_cap_vol_ouom, well_air_drill.Ppdm_guid, well_air_drill.Remark, well_air_drill.Row_changed_by, well_air_drill.Row_changed_date, well_air_drill.Row_created_by, well_air_drill.Row_created_date, well_air_drill.Row_effective_date, well_air_drill.Row_expiry_date, well_air_drill.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellAirDrill(c *fiber.Ctx) error {
	var well_air_drill dto.Well_air_drill

	setDefaults(&well_air_drill)

	if err := json.Unmarshal(c.Body(), &well_air_drill); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_air_drill.Uwi = id

    if well_air_drill.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_air_drill set source = :1, air_drill_obs_no = :2, active_ind = :3, compressor_count = :4, contractor = :5, effective_date = :6, expiry_date = :7, max_compressor_cap_vol = :8, max_compressor_cap_vol_ouom = :9, ppdm_guid = :10, remark = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where uwi = :19")
	if err != nil {
		return err
	}

	well_air_drill.Row_changed_date = formatDate(well_air_drill.Row_changed_date)
	well_air_drill.Effective_date = formatDateString(well_air_drill.Effective_date)
	well_air_drill.Expiry_date = formatDateString(well_air_drill.Expiry_date)
	well_air_drill.Row_effective_date = formatDateString(well_air_drill.Row_effective_date)
	well_air_drill.Row_expiry_date = formatDateString(well_air_drill.Row_expiry_date)






	rows, err := stmt.Exec(well_air_drill.Source, well_air_drill.Air_drill_obs_no, well_air_drill.Active_ind, well_air_drill.Compressor_count, well_air_drill.Contractor, well_air_drill.Effective_date, well_air_drill.Expiry_date, well_air_drill.Max_compressor_cap_vol, well_air_drill.Max_compressor_cap_vol_ouom, well_air_drill.Ppdm_guid, well_air_drill.Remark, well_air_drill.Row_changed_by, well_air_drill.Row_changed_date, well_air_drill.Row_created_by, well_air_drill.Row_effective_date, well_air_drill.Row_expiry_date, well_air_drill.Row_quality, well_air_drill.Uwi)
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

func PatchWellAirDrill(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_air_drill set "
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

func DeleteWellAirDrill(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_air_drill dto.Well_air_drill
	well_air_drill.Uwi = id

	stmt, err := db.Prepare("delete from well_air_drill where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_air_drill.Uwi)
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


