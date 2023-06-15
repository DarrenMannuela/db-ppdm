package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelTexas(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_texas")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_texas

	for rows.Next() {
		var sp_parcel_texas dto.Sp_parcel_texas
		if err := rows.Scan(&sp_parcel_texas.Parcel_texas_id, &sp_parcel_texas.Abstract_num, &sp_parcel_texas.Active_ind, &sp_parcel_texas.Area_id, &sp_parcel_texas.Area_type, &sp_parcel_texas.Block_fraction, &sp_parcel_texas.Coord_system_id, &sp_parcel_texas.Description, &sp_parcel_texas.Effective_date, &sp_parcel_texas.Expiry_date, &sp_parcel_texas.Labor, &sp_parcel_texas.League, &sp_parcel_texas.Ns_direction, &sp_parcel_texas.Porcion_num, &sp_parcel_texas.Porcion_survey_name, &sp_parcel_texas.Ppdm_guid, &sp_parcel_texas.Reference_plan_num, &sp_parcel_texas.Remark, &sp_parcel_texas.Section_fraction, &sp_parcel_texas.Source, &sp_parcel_texas.Spatial_description_id, &sp_parcel_texas.Spatial_obs_no, &sp_parcel_texas.Spot_code, &sp_parcel_texas.Texas_block, &sp_parcel_texas.Texas_lot, &sp_parcel_texas.Texas_section, &sp_parcel_texas.Texas_share, &sp_parcel_texas.Texas_subdivision, &sp_parcel_texas.Texas_survey, &sp_parcel_texas.Texas_township, &sp_parcel_texas.Row_changed_by, &sp_parcel_texas.Row_changed_date, &sp_parcel_texas.Row_created_by, &sp_parcel_texas.Row_created_date, &sp_parcel_texas.Row_effective_date, &sp_parcel_texas.Row_expiry_date, &sp_parcel_texas.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_texas to result
		result = append(result, sp_parcel_texas)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelTexas(c *fiber.Ctx) error {
	var sp_parcel_texas dto.Sp_parcel_texas

	setDefaults(&sp_parcel_texas)

	if err := json.Unmarshal(c.Body(), &sp_parcel_texas); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_texas values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	sp_parcel_texas.Row_created_date = formatDate(sp_parcel_texas.Row_created_date)
	sp_parcel_texas.Row_changed_date = formatDate(sp_parcel_texas.Row_changed_date)
	sp_parcel_texas.Effective_date = formatDateString(sp_parcel_texas.Effective_date)
	sp_parcel_texas.Expiry_date = formatDateString(sp_parcel_texas.Expiry_date)
	sp_parcel_texas.Row_effective_date = formatDateString(sp_parcel_texas.Row_effective_date)
	sp_parcel_texas.Row_expiry_date = formatDateString(sp_parcel_texas.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_texas.Parcel_texas_id, sp_parcel_texas.Abstract_num, sp_parcel_texas.Active_ind, sp_parcel_texas.Area_id, sp_parcel_texas.Area_type, sp_parcel_texas.Block_fraction, sp_parcel_texas.Coord_system_id, sp_parcel_texas.Description, sp_parcel_texas.Effective_date, sp_parcel_texas.Expiry_date, sp_parcel_texas.Labor, sp_parcel_texas.League, sp_parcel_texas.Ns_direction, sp_parcel_texas.Porcion_num, sp_parcel_texas.Porcion_survey_name, sp_parcel_texas.Ppdm_guid, sp_parcel_texas.Reference_plan_num, sp_parcel_texas.Remark, sp_parcel_texas.Section_fraction, sp_parcel_texas.Source, sp_parcel_texas.Spatial_description_id, sp_parcel_texas.Spatial_obs_no, sp_parcel_texas.Spot_code, sp_parcel_texas.Texas_block, sp_parcel_texas.Texas_lot, sp_parcel_texas.Texas_section, sp_parcel_texas.Texas_share, sp_parcel_texas.Texas_subdivision, sp_parcel_texas.Texas_survey, sp_parcel_texas.Texas_township, sp_parcel_texas.Row_changed_by, sp_parcel_texas.Row_changed_date, sp_parcel_texas.Row_created_by, sp_parcel_texas.Row_created_date, sp_parcel_texas.Row_effective_date, sp_parcel_texas.Row_expiry_date, sp_parcel_texas.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelTexas(c *fiber.Ctx) error {
	var sp_parcel_texas dto.Sp_parcel_texas

	setDefaults(&sp_parcel_texas)

	if err := json.Unmarshal(c.Body(), &sp_parcel_texas); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_texas.Parcel_texas_id = id

    if sp_parcel_texas.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_texas set abstract_num = :1, active_ind = :2, area_id = :3, area_type = :4, block_fraction = :5, coord_system_id = :6, description = :7, effective_date = :8, expiry_date = :9, labor = :10, league = :11, ns_direction = :12, porcion_num = :13, porcion_survey_name = :14, ppdm_guid = :15, reference_plan_num = :16, remark = :17, section_fraction = :18, source = :19, spatial_description_id = :20, spatial_obs_no = :21, spot_code = :22, texas_block = :23, texas_lot = :24, texas_section = :25, texas_share = :26, texas_subdivision = :27, texas_survey = :28, texas_township = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where parcel_texas_id = :37")
	if err != nil {
		return err
	}

	sp_parcel_texas.Row_changed_date = formatDate(sp_parcel_texas.Row_changed_date)
	sp_parcel_texas.Effective_date = formatDateString(sp_parcel_texas.Effective_date)
	sp_parcel_texas.Expiry_date = formatDateString(sp_parcel_texas.Expiry_date)
	sp_parcel_texas.Row_effective_date = formatDateString(sp_parcel_texas.Row_effective_date)
	sp_parcel_texas.Row_expiry_date = formatDateString(sp_parcel_texas.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_texas.Abstract_num, sp_parcel_texas.Active_ind, sp_parcel_texas.Area_id, sp_parcel_texas.Area_type, sp_parcel_texas.Block_fraction, sp_parcel_texas.Coord_system_id, sp_parcel_texas.Description, sp_parcel_texas.Effective_date, sp_parcel_texas.Expiry_date, sp_parcel_texas.Labor, sp_parcel_texas.League, sp_parcel_texas.Ns_direction, sp_parcel_texas.Porcion_num, sp_parcel_texas.Porcion_survey_name, sp_parcel_texas.Ppdm_guid, sp_parcel_texas.Reference_plan_num, sp_parcel_texas.Remark, sp_parcel_texas.Section_fraction, sp_parcel_texas.Source, sp_parcel_texas.Spatial_description_id, sp_parcel_texas.Spatial_obs_no, sp_parcel_texas.Spot_code, sp_parcel_texas.Texas_block, sp_parcel_texas.Texas_lot, sp_parcel_texas.Texas_section, sp_parcel_texas.Texas_share, sp_parcel_texas.Texas_subdivision, sp_parcel_texas.Texas_survey, sp_parcel_texas.Texas_township, sp_parcel_texas.Row_changed_by, sp_parcel_texas.Row_changed_date, sp_parcel_texas.Row_created_by, sp_parcel_texas.Row_effective_date, sp_parcel_texas.Row_expiry_date, sp_parcel_texas.Row_quality, sp_parcel_texas.Parcel_texas_id)
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

func PatchSpParcelTexas(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_texas set "
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
	query += " where parcel_texas_id = :id"

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

func DeleteSpParcelTexas(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_texas dto.Sp_parcel_texas
	sp_parcel_texas.Parcel_texas_id = id

	stmt, err := db.Prepare("delete from sp_parcel_texas where parcel_texas_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_texas.Parcel_texas_id)
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


