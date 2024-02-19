import paho.mqtt.client as mqtt
import time
from faker import Faker


def connect_client():
    client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
    client.connect("localhost", 1891, 60)
    return client


# Função para gerar dados do sensor MiCS-6814 simulados
def generate_sensor_data():
    fake = Faker()
    sensor_outputs = {
        "CO_ppm": fake.pyfloat(min_value=1, max_value=1000, right_digits=2),
        "NO2_ppm": fake.pyfloat(min_value=0.05, max_value=10, right_digits=2),
        "NH3_ppm": fake.pyfloat(min_value=1, max_value=300, right_digits=2),
    }
    return sensor_outputs


def main():
    client = connect_client()
    try:
        while True:
            sensor_data = generate_sensor_data()
            client.publish("sensor/mics6814", str(sensor_data))
            print(f"[SENSOR] Publishing message: {sensor_data}")
            time.sleep(1)
    except KeyboardInterrupt:
        print("[SENSOR] Publisher interrupted")
    except Exception as err:
        print(f"[SENSOR] Catch an exception: {str(err)}")
    finally:
        client.disconnect()


if __name__ == "__main__":
    main()
