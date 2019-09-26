package feast.serving.configuration;

import feast.core.CoreServiceProto.GetStoresRequest;
import feast.core.CoreServiceProto.GetStoresRequest.Filter;
import feast.core.CoreServiceProto.GetStoresResponse;
import feast.core.StoreProto.Store;
import feast.core.StoreProto.Store.RedisConfig;
import feast.core.StoreProto.Store.StoreType;
import feast.serving.service.BigQueryServingService;
import feast.serving.service.RedisServingService;
import feast.serving.service.ServingService;
import feast.serving.service.SpecService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Slf4j
@Configuration
public class ServingServiceConfig {
  @Value("${feast.store.name}")
  private String feastStoreName;

  @Bean
  public ServingService servingService(SpecService specService) {
    GetStoresResponse storesResponse =
        specService.getStores(
            GetStoresRequest.newBuilder()
                .setFilter(Filter.newBuilder().setName(feastStoreName).build())
                .build());
    if (storesResponse.getStoreCount() < 1) {
      throw new IllegalArgumentException(
          String.format(
              "Cannot resolve Store from store name '%s'. Ensure the store name exists in Feast.",
              feastStoreName));
    }
    assert storesResponse.getStoreCount() == 1;
    Store store = storesResponse.getStore(0);
    StoreType storeType = store.getType();
    ServingService servingService = null;

    switch (storeType) {
      case REDIS:
        RedisConfig redisConfig = store.getRedisConfig();
        servingService = new RedisServingService(redisConfig.getHost(), redisConfig.getPort());
        break;
      case BIGQUERY:
        servingService = new BigQueryServingService();
        break;
      case CASSANDRA:
      case UNRECOGNIZED:
      case INVALID:
        throw new IllegalArgumentException(
            String.format(
                "Unsupported store type '%s' for store name '%s'", storeType, feastStoreName));
    }

    return servingService;
  }
}