# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'AliasRoutingStrategy',
    'BuildStorageLocation',
    'FleetEc2InboundPermission',
    'FleetResourceCreationLimitPolicy',
    'FleetRuntimeConfiguration',
    'FleetRuntimeConfigurationServerProcess',
    'GameSessionQueuePlayerLatencyPolicy',
]

@pulumi.output_type
class AliasRoutingStrategy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fleetId":
            suggest = "fleet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AliasRoutingStrategy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AliasRoutingStrategy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AliasRoutingStrategy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 fleet_id: Optional[str] = None,
                 message: Optional[str] = None):
        """
        :param str type: Type of routing strategyE.g., `SIMPLE` or `TERMINAL`
        :param str fleet_id: ID of the Gamelift Fleet to point the alias to.
        :param str message: Message text to be used with the `TERMINAL` routing strategy.
        """
        pulumi.set(__self__, "type", type)
        if fleet_id is not None:
            pulumi.set(__self__, "fleet_id", fleet_id)
        if message is not None:
            pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of routing strategyE.g., `SIMPLE` or `TERMINAL`
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="fleetId")
    def fleet_id(self) -> Optional[str]:
        """
        ID of the Gamelift Fleet to point the alias to.
        """
        return pulumi.get(self, "fleet_id")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        Message text to be used with the `TERMINAL` routing strategy.
        """
        return pulumi.get(self, "message")


@pulumi.output_type
class BuildStorageLocation(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "roleArn":
            suggest = "role_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BuildStorageLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BuildStorageLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BuildStorageLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bucket: str,
                 key: str,
                 role_arn: str):
        """
        :param str bucket: Name of your S3 bucket.
        :param str key: Name of the zip file containing your build files.
        :param str role_arn: ARN of the access role that allows Amazon GameLift to access your S3 bucket.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def bucket(self) -> str:
        """
        Name of your S3 bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Name of the zip file containing your build files.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> str:
        """
        ARN of the access role that allows Amazon GameLift to access your S3 bucket.
        """
        return pulumi.get(self, "role_arn")


@pulumi.output_type
class FleetEc2InboundPermission(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fromPort":
            suggest = "from_port"
        elif key == "ipRange":
            suggest = "ip_range"
        elif key == "toPort":
            suggest = "to_port"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FleetEc2InboundPermission. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FleetEc2InboundPermission.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FleetEc2InboundPermission.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 from_port: int,
                 ip_range: str,
                 protocol: str,
                 to_port: int):
        """
        :param int from_port: Starting value for a range of allowed port numbers.
        :param str ip_range: Range of allowed IP addresses expressed in CIDR notationE.g., `000.000.000.000/[subnet mask]` or `0.0.0.0/[subnet mask]`.
        :param str protocol: Network communication protocol used by the fleetE.g., `TCP` or `UDP`
        :param int to_port: Ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than `from_port`.
        """
        pulumi.set(__self__, "from_port", from_port)
        pulumi.set(__self__, "ip_range", ip_range)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> int:
        """
        Starting value for a range of allowed port numbers.
        """
        return pulumi.get(self, "from_port")

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> str:
        """
        Range of allowed IP addresses expressed in CIDR notationE.g., `000.000.000.000/[subnet mask]` or `0.0.0.0/[subnet mask]`.
        """
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Network communication protocol used by the fleetE.g., `TCP` or `UDP`
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> int:
        """
        Ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than `from_port`.
        """
        return pulumi.get(self, "to_port")


@pulumi.output_type
class FleetResourceCreationLimitPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "newGameSessionsPerCreator":
            suggest = "new_game_sessions_per_creator"
        elif key == "policyPeriodInMinutes":
            suggest = "policy_period_in_minutes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FleetResourceCreationLimitPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FleetResourceCreationLimitPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FleetResourceCreationLimitPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 new_game_sessions_per_creator: Optional[int] = None,
                 policy_period_in_minutes: Optional[int] = None):
        """
        :param int new_game_sessions_per_creator: Maximum number of game sessions that an individual can create during the policy period.
        :param int policy_period_in_minutes: Time span used in evaluating the resource creation limit policy.
        """
        if new_game_sessions_per_creator is not None:
            pulumi.set(__self__, "new_game_sessions_per_creator", new_game_sessions_per_creator)
        if policy_period_in_minutes is not None:
            pulumi.set(__self__, "policy_period_in_minutes", policy_period_in_minutes)

    @property
    @pulumi.getter(name="newGameSessionsPerCreator")
    def new_game_sessions_per_creator(self) -> Optional[int]:
        """
        Maximum number of game sessions that an individual can create during the policy period.
        """
        return pulumi.get(self, "new_game_sessions_per_creator")

    @property
    @pulumi.getter(name="policyPeriodInMinutes")
    def policy_period_in_minutes(self) -> Optional[int]:
        """
        Time span used in evaluating the resource creation limit policy.
        """
        return pulumi.get(self, "policy_period_in_minutes")


@pulumi.output_type
class FleetRuntimeConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "gameSessionActivationTimeoutSeconds":
            suggest = "game_session_activation_timeout_seconds"
        elif key == "maxConcurrentGameSessionActivations":
            suggest = "max_concurrent_game_session_activations"
        elif key == "serverProcesses":
            suggest = "server_processes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FleetRuntimeConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FleetRuntimeConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FleetRuntimeConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 game_session_activation_timeout_seconds: Optional[int] = None,
                 max_concurrent_game_session_activations: Optional[int] = None,
                 server_processes: Optional[Sequence['outputs.FleetRuntimeConfigurationServerProcess']] = None):
        """
        :param int game_session_activation_timeout_seconds: Maximum amount of time (in seconds) that a game session can remain in status `ACTIVATING`.
        :param int max_concurrent_game_session_activations: Maximum number of game sessions with status `ACTIVATING` to allow on an instance simultaneously.
        :param Sequence['FleetRuntimeConfigurationServerProcessArgs'] server_processes: Collection of server process configurations that describe which server processes to run on each instance in a fleet. See below.
        """
        if game_session_activation_timeout_seconds is not None:
            pulumi.set(__self__, "game_session_activation_timeout_seconds", game_session_activation_timeout_seconds)
        if max_concurrent_game_session_activations is not None:
            pulumi.set(__self__, "max_concurrent_game_session_activations", max_concurrent_game_session_activations)
        if server_processes is not None:
            pulumi.set(__self__, "server_processes", server_processes)

    @property
    @pulumi.getter(name="gameSessionActivationTimeoutSeconds")
    def game_session_activation_timeout_seconds(self) -> Optional[int]:
        """
        Maximum amount of time (in seconds) that a game session can remain in status `ACTIVATING`.
        """
        return pulumi.get(self, "game_session_activation_timeout_seconds")

    @property
    @pulumi.getter(name="maxConcurrentGameSessionActivations")
    def max_concurrent_game_session_activations(self) -> Optional[int]:
        """
        Maximum number of game sessions with status `ACTIVATING` to allow on an instance simultaneously.
        """
        return pulumi.get(self, "max_concurrent_game_session_activations")

    @property
    @pulumi.getter(name="serverProcesses")
    def server_processes(self) -> Optional[Sequence['outputs.FleetRuntimeConfigurationServerProcess']]:
        """
        Collection of server process configurations that describe which server processes to run on each instance in a fleet. See below.
        """
        return pulumi.get(self, "server_processes")


@pulumi.output_type
class FleetRuntimeConfigurationServerProcess(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "concurrentExecutions":
            suggest = "concurrent_executions"
        elif key == "launchPath":
            suggest = "launch_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FleetRuntimeConfigurationServerProcess. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FleetRuntimeConfigurationServerProcess.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FleetRuntimeConfigurationServerProcess.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 concurrent_executions: int,
                 launch_path: str,
                 parameters: Optional[str] = None):
        """
        :param int concurrent_executions: Number of server processes using this configuration to run concurrently on an instance.
        :param str launch_path: Location of the server executable in a game build. All game builds are installed on instances at the root : for Windows instances `C:\game`, and for Linux instances `/local/game`.
        :param str parameters: Optional list of parameters to pass to the server executable on launch.
        """
        pulumi.set(__self__, "concurrent_executions", concurrent_executions)
        pulumi.set(__self__, "launch_path", launch_path)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="concurrentExecutions")
    def concurrent_executions(self) -> int:
        """
        Number of server processes using this configuration to run concurrently on an instance.
        """
        return pulumi.get(self, "concurrent_executions")

    @property
    @pulumi.getter(name="launchPath")
    def launch_path(self) -> str:
        """
        Location of the server executable in a game build. All game builds are installed on instances at the root : for Windows instances `C:\game`, and for Linux instances `/local/game`.
        """
        return pulumi.get(self, "launch_path")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[str]:
        """
        Optional list of parameters to pass to the server executable on launch.
        """
        return pulumi.get(self, "parameters")


@pulumi.output_type
class GameSessionQueuePlayerLatencyPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maximumIndividualPlayerLatencyMilliseconds":
            suggest = "maximum_individual_player_latency_milliseconds"
        elif key == "policyDurationSeconds":
            suggest = "policy_duration_seconds"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GameSessionQueuePlayerLatencyPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GameSessionQueuePlayerLatencyPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GameSessionQueuePlayerLatencyPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 maximum_individual_player_latency_milliseconds: int,
                 policy_duration_seconds: Optional[int] = None):
        """
        :param int maximum_individual_player_latency_milliseconds: Maximum latency value that is allowed for any player.
        :param int policy_duration_seconds: Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
        """
        pulumi.set(__self__, "maximum_individual_player_latency_milliseconds", maximum_individual_player_latency_milliseconds)
        if policy_duration_seconds is not None:
            pulumi.set(__self__, "policy_duration_seconds", policy_duration_seconds)

    @property
    @pulumi.getter(name="maximumIndividualPlayerLatencyMilliseconds")
    def maximum_individual_player_latency_milliseconds(self) -> int:
        """
        Maximum latency value that is allowed for any player.
        """
        return pulumi.get(self, "maximum_individual_player_latency_milliseconds")

    @property
    @pulumi.getter(name="policyDurationSeconds")
    def policy_duration_seconds(self) -> Optional[int]:
        """
        Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
        """
        return pulumi.get(self, "policy_duration_seconds")


