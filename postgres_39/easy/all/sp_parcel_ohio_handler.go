package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelOhio(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_ohio")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_ohio

	for rows.Next() {
		var sp_parcel_ohio dto.Sp_parcel_ohio
		if err := rows.Scan(&sp_parcel_ohio.Parcel_ohio_id, &sp_parcel_ohio.Active_ind, &sp_parcel_ohio.Area_id, &sp_parcel_ohio.Area_type, &sp_parcel_ohio.Coord_system_id, &sp_parcel_ohio.Description, &sp_parcel_ohio.Effective_date, &sp_parcel_ohio.Expiry_date, &sp_parcel_ohio.Jurisdiction, &sp_parcel_ohio.Map_quad_min, &sp_parcel_ohio.Map_quad_name, &sp_parcel_ohio.Ohio_allotment, &sp_parcel_ohio.Ohio_division, &sp_parcel_ohio.Ohio_fraction, &sp_parcel_ohio.Ohio_land_subdivision_name, &sp_parcel_ohio.Ohio_other_subdivision, &sp_parcel_ohio.Ohio_quarter_township, &sp_parcel_ohio.Ohio_range, &sp_parcel_ohio.Ohio_range_dir, &sp_parcel_ohio.Ohio_township, &sp_parcel_ohio.Ohio_township_dir, &sp_parcel_ohio.Ohio_township_name, &sp_parcel_ohio.Ohio_tract, &sp_parcel_ohio.Ohio_twp_lot_code, &sp_parcel_ohio.Ohio_twp_section_code, &sp_parcel_ohio.Ppdm_guid, &sp_parcel_ohio.Principal_meridian, &sp_parcel_ohio.Reference_plan_num, &sp_parcel_ohio.Remark, &sp_parcel_ohio.Source, &sp_parcel_ohio.Spatial_description_id, &sp_parcel_ohio.Spatial_obs_no, &sp_parcel_ohio.Spot_code, &sp_parcel_ohio.Row_changed_by, &sp_parcel_ohio.Row_changed_date, &sp_parcel_ohio.Row_created_by, &sp_parcel_ohio.Row_created_date, &sp_parcel_ohio.Row_effective_date, &sp_parcel_ohio.Row_expiry_date, &sp_parcel_ohio.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_ohio to result
		result = append(result, sp_parcel_ohio)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelOhio(c *fiber.Ctx) error {
	var sp_parcel_ohio dto.Sp_parcel_ohio

	setDefaults(&sp_parcel_ohio)

	if err := json.Unmarshal(c.Body(), &sp_parcel_ohio); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_ohio values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	sp_parcel_ohio.Row_created_date = formatDate(sp_parcel_ohio.Row_created_date)
	sp_parcel_ohio.Row_changed_date = formatDate(sp_parcel_ohio.Row_changed_date)
	sp_parcel_ohio.Effective_date = formatDateString(sp_parcel_ohio.Effective_date)
	sp_parcel_ohio.Expiry_date = formatDateString(sp_parcel_ohio.Expiry_date)
	sp_parcel_ohio.Row_effective_date = formatDateString(sp_parcel_ohio.Row_effective_date)
	sp_parcel_ohio.Row_expiry_date = formatDateString(sp_parcel_ohio.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_ohio.Parcel_ohio_id, sp_parcel_ohio.Active_ind, sp_parcel_ohio.Area_id, sp_parcel_ohio.Area_type, sp_parcel_ohio.Coord_system_id, sp_parcel_ohio.Description, sp_parcel_ohio.Effective_date, sp_parcel_ohio.Expiry_date, sp_parcel_ohio.Jurisdiction, sp_parcel_ohio.Map_quad_min, sp_parcel_ohio.Map_quad_name, sp_parcel_ohio.Ohio_allotment, sp_parcel_ohio.Ohio_division, sp_parcel_ohio.Ohio_fraction, sp_parcel_ohio.Ohio_land_subdivision_name, sp_parcel_ohio.Ohio_other_subdivision, sp_parcel_ohio.Ohio_quarter_township, sp_parcel_ohio.Ohio_range, sp_parcel_ohio.Ohio_range_dir, sp_parcel_ohio.Ohio_township, sp_parcel_ohio.Ohio_township_dir, sp_parcel_ohio.Ohio_township_name, sp_parcel_ohio.Ohio_tract, sp_parcel_ohio.Ohio_twp_lot_code, sp_parcel_ohio.Ohio_twp_section_code, sp_parcel_ohio.Ppdm_guid, sp_parcel_ohio.Principal_meridian, sp_parcel_ohio.Reference_plan_num, sp_parcel_ohio.Remark, sp_parcel_ohio.Source, sp_parcel_ohio.Spatial_description_id, sp_parcel_ohio.Spatial_obs_no, sp_parcel_ohio.Spot_code, sp_parcel_ohio.Row_changed_by, sp_parcel_ohio.Row_changed_date, sp_parcel_ohio.Row_created_by, sp_parcel_ohio.Row_created_date, sp_parcel_ohio.Row_effective_date, sp_parcel_ohio.Row_expiry_date, sp_parcel_ohio.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelOhio(c *fiber.Ctx) error {
	var sp_parcel_ohio dto.Sp_parcel_ohio

	setDefaults(&sp_parcel_ohio)

	if err := json.Unmarshal(c.Body(), &sp_parcel_ohio); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_ohio.Parcel_ohio_id = id

    if sp_parcel_ohio.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_ohio set active_ind = :1, area_id = :2, area_type = :3, coord_system_id = :4, description = :5, effective_date = :6, expiry_date = :7, jurisdiction = :8, map_quad_min = :9, map_quad_name = :10, ohio_allotment = :11, ohio_division = :12, ohio_fraction = :13, ohio_land_subdivision_name = :14, ohio_other_subdivision = :15, ohio_quarter_township = :16, ohio_range = :17, ohio_range_dir = :18, ohio_township = :19, ohio_township_dir = :20, ohio_township_name = :21, ohio_tract = :22, ohio_twp_lot_code = :23, ohio_twp_section_code = :24, ppdm_guid = :25, principal_meridian = :26, reference_plan_num = :27, remark = :28, source = :29, spatial_description_id = :30, spatial_obs_no = :31, spot_code = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where parcel_ohio_id = :40")
	if err != nil {
		return err
	}

	sp_parcel_ohio.Row_changed_date = formatDate(sp_parcel_ohio.Row_changed_date)
	sp_parcel_ohio.Effective_date = formatDateString(sp_parcel_ohio.Effective_date)
	sp_parcel_ohio.Expiry_date = formatDateString(sp_parcel_ohio.Expiry_date)
	sp_parcel_ohio.Row_effective_date = formatDateString(sp_parcel_ohio.Row_effective_date)
	sp_parcel_ohio.Row_expiry_date = formatDateString(sp_parcel_ohio.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_ohio.Active_ind, sp_parcel_ohio.Area_id, sp_parcel_ohio.Area_type, sp_parcel_ohio.Coord_system_id, sp_parcel_ohio.Description, sp_parcel_ohio.Effective_date, sp_parcel_ohio.Expiry_date, sp_parcel_ohio.Jurisdiction, sp_parcel_ohio.Map_quad_min, sp_parcel_ohio.Map_quad_name, sp_parcel_ohio.Ohio_allotment, sp_parcel_ohio.Ohio_division, sp_parcel_ohio.Ohio_fraction, sp_parcel_ohio.Ohio_land_subdivision_name, sp_parcel_ohio.Ohio_other_subdivision, sp_parcel_ohio.Ohio_quarter_township, sp_parcel_ohio.Ohio_range, sp_parcel_ohio.Ohio_range_dir, sp_parcel_ohio.Ohio_township, sp_parcel_ohio.Ohio_township_dir, sp_parcel_ohio.Ohio_township_name, sp_parcel_ohio.Ohio_tract, sp_parcel_ohio.Ohio_twp_lot_code, sp_parcel_ohio.Ohio_twp_section_code, sp_parcel_ohio.Ppdm_guid, sp_parcel_ohio.Principal_meridian, sp_parcel_ohio.Reference_plan_num, sp_parcel_ohio.Remark, sp_parcel_ohio.Source, sp_parcel_ohio.Spatial_description_id, sp_parcel_ohio.Spatial_obs_no, sp_parcel_ohio.Spot_code, sp_parcel_ohio.Row_changed_by, sp_parcel_ohio.Row_changed_date, sp_parcel_ohio.Row_created_by, sp_parcel_ohio.Row_effective_date, sp_parcel_ohio.Row_expiry_date, sp_parcel_ohio.Row_quality, sp_parcel_ohio.Parcel_ohio_id)
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

func PatchSpParcelOhio(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_ohio set "
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
	query += " where parcel_ohio_id = :id"

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

func DeleteSpParcelOhio(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_ohio dto.Sp_parcel_ohio
	sp_parcel_ohio.Parcel_ohio_id = id

	stmt, err := db.Prepare("delete from sp_parcel_ohio where parcel_ohio_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_ohio.Parcel_ohio_id)
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


