package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelNts(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_nts")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_nts

	for rows.Next() {
		var sp_parcel_nts dto.Sp_parcel_nts
		if err := rows.Scan(&sp_parcel_nts.Parcel_nts_id, &sp_parcel_nts.Active_ind, &sp_parcel_nts.Area_id, &sp_parcel_nts.Area_type, &sp_parcel_nts.Block, &sp_parcel_nts.Coord_system_id, &sp_parcel_nts.Description, &sp_parcel_nts.Effective_date, &sp_parcel_nts.Expiry_date, &sp_parcel_nts.Letter_quadrangle, &sp_parcel_nts.Ppdm_guid, &sp_parcel_nts.Primary_quadrangle, &sp_parcel_nts.Quarter_unit, &sp_parcel_nts.Reference_plan_num, &sp_parcel_nts.Remark, &sp_parcel_nts.Sixteenth, &sp_parcel_nts.Source, &sp_parcel_nts.Spatial_description_id, &sp_parcel_nts.Spatial_obs_no, &sp_parcel_nts.Unit, &sp_parcel_nts.Row_changed_by, &sp_parcel_nts.Row_changed_date, &sp_parcel_nts.Row_created_by, &sp_parcel_nts.Row_created_date, &sp_parcel_nts.Row_effective_date, &sp_parcel_nts.Row_expiry_date, &sp_parcel_nts.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_nts to result
		result = append(result, sp_parcel_nts)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelNts(c *fiber.Ctx) error {
	var sp_parcel_nts dto.Sp_parcel_nts

	setDefaults(&sp_parcel_nts)

	if err := json.Unmarshal(c.Body(), &sp_parcel_nts); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_nts values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	sp_parcel_nts.Row_created_date = formatDate(sp_parcel_nts.Row_created_date)
	sp_parcel_nts.Row_changed_date = formatDate(sp_parcel_nts.Row_changed_date)
	sp_parcel_nts.Effective_date = formatDateString(sp_parcel_nts.Effective_date)
	sp_parcel_nts.Expiry_date = formatDateString(sp_parcel_nts.Expiry_date)
	sp_parcel_nts.Row_effective_date = formatDateString(sp_parcel_nts.Row_effective_date)
	sp_parcel_nts.Row_expiry_date = formatDateString(sp_parcel_nts.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_nts.Parcel_nts_id, sp_parcel_nts.Active_ind, sp_parcel_nts.Area_id, sp_parcel_nts.Area_type, sp_parcel_nts.Block, sp_parcel_nts.Coord_system_id, sp_parcel_nts.Description, sp_parcel_nts.Effective_date, sp_parcel_nts.Expiry_date, sp_parcel_nts.Letter_quadrangle, sp_parcel_nts.Ppdm_guid, sp_parcel_nts.Primary_quadrangle, sp_parcel_nts.Quarter_unit, sp_parcel_nts.Reference_plan_num, sp_parcel_nts.Remark, sp_parcel_nts.Sixteenth, sp_parcel_nts.Source, sp_parcel_nts.Spatial_description_id, sp_parcel_nts.Spatial_obs_no, sp_parcel_nts.Unit, sp_parcel_nts.Row_changed_by, sp_parcel_nts.Row_changed_date, sp_parcel_nts.Row_created_by, sp_parcel_nts.Row_created_date, sp_parcel_nts.Row_effective_date, sp_parcel_nts.Row_expiry_date, sp_parcel_nts.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelNts(c *fiber.Ctx) error {
	var sp_parcel_nts dto.Sp_parcel_nts

	setDefaults(&sp_parcel_nts)

	if err := json.Unmarshal(c.Body(), &sp_parcel_nts); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_nts.Parcel_nts_id = id

    if sp_parcel_nts.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_nts set active_ind = :1, area_id = :2, area_type = :3, block = :4, coord_system_id = :5, description = :6, effective_date = :7, expiry_date = :8, letter_quadrangle = :9, ppdm_guid = :10, primary_quadrangle = :11, quarter_unit = :12, reference_plan_num = :13, remark = :14, sixteenth = :15, source = :16, spatial_description_id = :17, spatial_obs_no = :18, unit = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where parcel_nts_id = :27")
	if err != nil {
		return err
	}

	sp_parcel_nts.Row_changed_date = formatDate(sp_parcel_nts.Row_changed_date)
	sp_parcel_nts.Effective_date = formatDateString(sp_parcel_nts.Effective_date)
	sp_parcel_nts.Expiry_date = formatDateString(sp_parcel_nts.Expiry_date)
	sp_parcel_nts.Row_effective_date = formatDateString(sp_parcel_nts.Row_effective_date)
	sp_parcel_nts.Row_expiry_date = formatDateString(sp_parcel_nts.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_nts.Active_ind, sp_parcel_nts.Area_id, sp_parcel_nts.Area_type, sp_parcel_nts.Block, sp_parcel_nts.Coord_system_id, sp_parcel_nts.Description, sp_parcel_nts.Effective_date, sp_parcel_nts.Expiry_date, sp_parcel_nts.Letter_quadrangle, sp_parcel_nts.Ppdm_guid, sp_parcel_nts.Primary_quadrangle, sp_parcel_nts.Quarter_unit, sp_parcel_nts.Reference_plan_num, sp_parcel_nts.Remark, sp_parcel_nts.Sixteenth, sp_parcel_nts.Source, sp_parcel_nts.Spatial_description_id, sp_parcel_nts.Spatial_obs_no, sp_parcel_nts.Unit, sp_parcel_nts.Row_changed_by, sp_parcel_nts.Row_changed_date, sp_parcel_nts.Row_created_by, sp_parcel_nts.Row_effective_date, sp_parcel_nts.Row_expiry_date, sp_parcel_nts.Row_quality, sp_parcel_nts.Parcel_nts_id)
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

func PatchSpParcelNts(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_nts set "
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
	query += " where parcel_nts_id = :id"

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

func DeleteSpParcelNts(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_nts dto.Sp_parcel_nts
	sp_parcel_nts.Parcel_nts_id = id

	stmt, err := db.Prepare("delete from sp_parcel_nts where parcel_nts_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_nts.Parcel_nts_id)
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


