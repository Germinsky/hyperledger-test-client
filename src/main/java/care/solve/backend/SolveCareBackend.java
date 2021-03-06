package care.solve.backend;

import care.solve.fabric.FabricSdk;
import care.solve.protocol.schedule.ScheduleProtocol;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.AutoConfigurationExcludeFilter;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.TypeExcludeFilter;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.FilterType;

@SpringBootApplication
@ComponentScan(basePackageClasses = {
        SolveCareBackend.class, FabricSdk.class, ScheduleProtocol.class
},
        excludeFilters = {
        @ComponentScan.Filter(type = FilterType.CUSTOM, classes = TypeExcludeFilter.class),
        @ComponentScan.Filter(type = FilterType.CUSTOM, classes = AutoConfigurationExcludeFilter.class) })
public class SolveCareBackend {

    public static void main(String[] args) {
        SpringApplication.run(SolveCareBackend.class, args);
    }
}
