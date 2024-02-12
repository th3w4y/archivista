-- Modify "attribute_assertions" table
ALTER TABLE "attribute_assertions" DROP CONSTRAINT "attribute_assertions_attribute_reports_attributes", ALTER COLUMN "attribute_report_attributes" SET NOT NULL, ADD CONSTRAINT "attribute_assertions_attribute_reports_attributes" FOREIGN KEY ("attribute_report_attributes") REFERENCES "attribute_reports" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
